package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
)

func OrderDetailsGet(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	var para models.OrderDetailsCondition

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

	if userPayload.AccountType != "owner" {
		response.Status = "403"
		response.Message = "You are not authorized for this request."
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}

	para.Id = Helper.StrToInt(r.URL.Query().Get("id"))               // take id for url
	para.CompanyId = Helper.StrToInt(r.URL.Query().Get("CompanyId")) // take company_id for url

	parameters := models.OrderDetails{}.GetFilterOrderDetails(para)
	result, err := models.OrderDetails{}.OrderDetailsGet(parameters)
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
		response.Message = "Returns all matching order details."
		response.Payload = result
	}
	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}
