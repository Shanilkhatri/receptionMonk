package controllers

import (
	"net/http"
	"reakgo/utility"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}

	// productId := Helper.StrToInt(r.URL.Query().Get("id"))

	isOk, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isOk || userDetails.AccountType != "" {
		response.Status = "403"
		response.Message = "Unauthorized access, you are not allowed to make this request!"
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	// now fetch product

}
