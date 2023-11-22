package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
)

func PutKycDetails(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var userStruct models.KycDetails
	err := utility.StrictParseDataFromJson(r, &userStruct)
	if err != nil {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	if userStruct.UserId > 0 && userStruct.DocPicName != "" && userStruct.DocName != "" {
		log.Println("userStruct: ", userStruct)
		tx := utility.Db.MustBegin()
		inserted := models.KycDetails{}.Putkyc(userStruct, tx)
		if inserted {
			err = tx.Commit()
			if err != nil {
				log.Println(err)
				tx.Rollback()
				response.Status = "400"
				response.Message = "Unable to update kyc at the moment! Please try again."
				utility.RenderJsonResponse(w, r, response, 400)
				return
			}
			response.Status = "200"
			response.Message = "Document upload successfully"
			utility.RenderJsonResponse(w, r, response, 200)
			return
		} else {
			tx.Rollback()
			response.Status = "400"
			response.Message = "Unable to update kyc at the moment! Please try again."
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
	} else {
		response.Status = "400"
		response.Message = "Please provide all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
}
func PostKycDetails(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var userStruct models.KycDetails
	err := utility.StrictParseDataFromJson(r, &userStruct)
	if err != nil {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	if userStruct.UserId > 0 && userStruct.DocPicName != "" && userStruct.DocName != "" {
		log.Println("userStruct: ", userStruct)
		tx := utility.Db.MustBegin()
		inserted, err := models.KycDetails{}.Postkyc(userStruct, tx)
		if inserted && err == nil {
			err = tx.Commit()
			if err != nil {
				log.Println(err)
				tx.Rollback()
				response.Status = "400"
				response.Message = "Unable to update kyc at the moment! Please try again."
				utility.RenderJsonResponse(w, r, response, 400)
				return
			}
			response.Status = "200"
			response.Message = "Document upload successfully"
			utility.RenderJsonResponse(w, r, response, 200)
			return
		} else {
			tx.Rollback()
			response.Status = "400"
			response.Message = "Unable to update kyc at the moment! Please try again."
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
	} else {
		response.Status = "400"
		response.Message = "Please provide all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
}
func GetKyc(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	isok, userDetailsType := utility.CheckTokenPayloadAndReturnUser(r)
	if !isok {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	kyc_data, err := models.KycDetails{}.Getkyc(userDetailsType.ID)
	if err != nil {
		response.Status = "400"
		response.Message = "Unable to get kyc at the moment! Please try again."
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	response.Status = "200"
	response.Message = "successfully getting the record."
	response.Payload = kyc_data
	utility.RenderJsonResponse(w, r, response, 200)
	return
}
