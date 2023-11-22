package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
)

func PutCompany(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var companyStruct models.Company
	err := utility.StrictParseDataFromJson(r, &companyStruct)
	log.Println("put companyStruct: ", companyStruct)
	if err != nil {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	ok, userDetails := utility.CheckTokenPayloadAndReturnUser(r)
	if !ok {
		utility.Logger(err)
		response.Status = "403"
		response.Message = "You cannot register the company because you are not an owner."
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	userType := userDetails.AccountType
	// userType := utility.SessionGet(r, "type")
	if userType == "" {
		userType = "guest"
	}
	if userType != "owner" {
		utility.Logger(err)
		response.Status = "403"
		response.Message = "You cannot register the company because you are not an owner."
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	tx := utility.Db.MustBegin()
	companyId, err1 := models.Company{}.PutCompany(companyStruct, tx)
	if err1 != nil {
		response.Status = "403"
		response.Message = "Unable to add company details at the moment! Please try again."
		isok, errString := utility.CheckSqlError(err, "") // dummy check
		if isok {
			log.Println(errString)
		}
		utility.Logger(err)
		tx.Rollback()
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	ok, err3 := models.Authentication{}.UpdateCompanyIdByEmail(userDetails.ID, companyId, tx)
	if err3 != nil {
		response.Status = "403"
		response.Message = "Unable to add company details at the moment! Please try again."
		isok, errString := utility.CheckSqlError(err, "") // dummy check
		if isok {
			log.Println(errString)
		}
		utility.Logger(err)
		tx.Rollback()
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	if ok {
		err = tx.Commit()
		if err != nil {
			log.Println(err)
			utility.Logger(err)
			tx.Rollback()
			response.Status = "400"
			response.Message = "Unable to add company details at the moment! Please try again."
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
	} else {
		tx.Rollback()
		response.Status = "400"
		response.Message = "Unable to add company details at the moment! Please try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}

	response.Status = "200"
	response.Message = "company details added successfully."
	utility.RenderJsonResponse(w, r, response, 200)
}

func PostCompany(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var companyStruct models.Company
	err := utility.StrictParseDataFromJson(r, &companyStruct)
	log.Println("companyStruct: ", companyStruct)
	if err != nil {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	ok, userDetails := utility.CheckTokenPayloadAndReturnUser(r)
	if !ok {
		utility.Logger(err)
		response.Status = "403"
		response.Message = "You cannot register the company because you are not an owner."
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	userType := userDetails.AccountType
	if userType == "" {
		userType = "guest"
	}
	if userType != "owner" {
		utility.Logger(err)
		response.Status = "403"
		response.Message = "You cannot register the company because you are not an owner."
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	tx := utility.Db.MustBegin()
	isupdate, err1 := models.Company{}.PostCompany(companyStruct, tx)
	if err1 != nil || !isupdate {
		response.Status = "403"
		response.Message = "Unable to add company details at the moment! Please try again."
		isok, errString := utility.CheckSqlError(err, "") // dummy check
		if isok {
			log.Println(errString)
		}
		utility.Logger(err)
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		utility.Logger(err)
		tx.Rollback()
		response.Status = "400"
		response.Message = "Unable to add company details at the moment! Please try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	response.Status = "200"
	response.Message = "company details updated successfully."
	utility.RenderJsonResponse(w, r, response, 200)
}

// isok, userDetails := utility.CheckTokenPayloadAndReturnUser(r)
// if !isok || userDetails.AccountType == "user" {
// 	response.Status = "403"
// 	response.Message = "You are not authorized to make this request."
// 	utility.RenderJsonResponse(w, r, response, 403)
// 	return
// }
// ok, err3 := models.Authentication{}.UpdateCompanyIdByEmail(userDetails.ID, companyId, tx)
// if err3 != nil {
// 	response.Status = "403"
// 	response.Message = "Unable to add company details at the moment! Please try again."
// 	isok, errString := utility.CheckSqlError(err, "") // dummy check
// 	if isok {
// 		log.Println(errString)
// 	}
// 	utility.Logger(err)
// 	utility.RenderJsonResponse(w, r, response, 400)
// 	return
// }
// if ok {
// } else {

// }
