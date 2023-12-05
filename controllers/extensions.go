package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"

	"github.com/go-sql-driver/mysql"
)

func PostExtension(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	var extensionStruct models.Extensions
	// var userPayload models.Users

	isOk, userPayload := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isOk {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	if userPayload.ID == 0 || userPayload.CompanyID == 0 {
		response.Status = "403"
		response.Message = "Unauthorized access, UserId or companyId doesn't match."
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}

	err := Helper.StrictParseDataFromJson(r, &extensionStruct)
	if err != nil {
		log.Println(err)
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	//check validation.
	boolType := ValidationCheck(extensionStruct)
	if boolType {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	tx := utility.Db.MustBegin()

	//Update data in table.
	boolValue, err := models.Extensions{}.PostExtension(extensionStruct, tx)

	if !boolValue || err != nil {
		log.Println(err)
		tx.Rollback()
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}

	txError := tx.Commit()
	if txError != nil {
		tx.Rollback()
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	} else {
		response.Status = "200"
		response.Message = "Extension updated successfully."
	}

	Helper.RenderJsonResponse(w, r, response, 200)
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

func PutExtension(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}

	var extensionStruct models.Extensions
	// var userPayload models.Users

	isOk, userPayload := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isOk {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	if userPayload.ID == 0 || userPayload.CompanyID == 0 {
		response.Status = "403"
		response.Message = "Unauthorized access, UserId or companyId doesn't match."
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}

	err := Helper.StrictParseDataFromJson(r, &extensionStruct)
	if err != nil {
		log.Println(err)
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}
	//check validation.
	boolType := ValidationCheck(extensionStruct)
	if boolType {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	tx := utility.Db.MustBegin()

	//put data in table.
	boolValue, err := models.Extensions{}.PutExtension(extensionStruct, tx)
	if boolValue || err != nil {
		log.Println(err)
		tx.Rollback()
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}

	txError := tx.Commit()
	if txError != nil {
		tx.Rollback()
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	} else {
		response.Status = "200"
		response.Message = "Extension added successfully."
	}

	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}

func DeleteExtension(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}

	extensionId := Helper.StrToInt(r.URL.Query().Get("id"))

	if extensionId <= 0 {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	boolType, err := models.Extensions{}.DeleteExtensionDetail(extensionId)
	if !boolType || err != nil {
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			// MySQL error code 1451 indicates a foreign key constraint
			if driverErr.Number == 1451 {
				response.Message = Helper.GetSqlErrorString(err)
			}
		}
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}
	response.Status = "200"
	response.Message = "Extension deleted successfully."
	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}

func GetExtensionData(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	var para models.ExtensionCondition

	// var userPayload models.Users

	isOk, userPayload := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isOk {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	if userPayload.ID == 0 || userPayload.CompanyID == 0 {
		response.Status = "403"
		response.Message = "Unauthorized access, UserId or companyId doesn't match."
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}

	para.AccountType = userPayload.AccountType

	para.Id = Helper.StrToInt(r.URL.Query().Get("id"))               // take id for url
	para.CompanyId = Helper.StrToInt(r.URL.Query().Get("companyId")) // take company_id for url

	if userPayload.ID != int64(para.CompanyId) {
		response.Status = "403"
		response.Message = "You are not authorized for this request."
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}

	if userPayload.ID != 0 && para.AccountType == "user" {
		para.Id = int(userPayload.ID)
	}

	parameters := models.Extensions{}.GetParaForFilterExtension(para)
	result, err := models.Extensions{}.GetExtensions(parameters)
	if err != nil {
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}

	if len(result) == 0 {
		response.Status = "200"
		response.Message = "No result were found for this search."
	} else {
		response.Status = "200"
		response.Message = "Returns all matching  Extensions."
		response.Payload = result // Set the extension data in the response payload
	}
	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}
