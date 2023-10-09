package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
)

func PostExtension(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponse{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	var extensionStruct models.Extensions
	var userPayload models.Users

	err := Utility.ReturnUserDetails(r, &userPayload)
	if err != nil {
		log.Println("====================")
		log.Println(err)
		log.Println("====================")
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Utility.RenderJsonResponse(w, r, response)
		return true
	}

	if userPayload.Id == 0 || userPayload.CompanyId == 0 {
		response.Status = "403"
		response.Message = "Unauthorized access, UserId or companyId doesn't match."
		Utility.RenderJsonResponse(w, r, response)
		return true
	}

	err = Utility.StrictParseDataFromJson(r, &extensionStruct)
	if err != nil {
		log.Println(err)
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Utility.RenderJsonResponse(w, r, response)
		return true
	}

	//check validation.
	boolType := ValidationCheck(extensionStruct)
	if boolType {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Utility.RenderJsonResponse(w, r, response)
		return true
	}

	tx := utility.Db.MustBegin()

	//Update data in table.
	boolValue, err := models.Extensions{}.PostExtension(extensionStruct, tx)
	log.Println(boolValue, err)

	if !boolValue || err != nil {
		log.Println(err)
		tx.Rollback()
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Utility.RenderJsonResponse(w, r, response)
		return true
	}

	txError := tx.Commit()
	if txError != nil {
		tx.Rollback()
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
	} else {
		response.Status = "200"
		response.Message = "Extension 1updated2 successfully."
	}

	Utility.RenderJsonResponse(w, r, response)
	return false
}

func ValidationCheck(extensionStruct models.Extensions) bool {
	switch {
	case extensionStruct.Extension == "":
		return true
	case extensionStruct.UserId <= 0:
		return true
	case extensionStruct.Department <= 0:
		return true
	case extensionStruct.SipServer == "":
		return true
	case extensionStruct.SipUserName == "":
		return true
	case extensionStruct.SipPassword == "":
		return true
	case extensionStruct.SipPort == "":
		return true
	default:
		return false
	}
}
