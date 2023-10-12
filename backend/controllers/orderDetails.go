package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
)

func OrderDetailsGet(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponse{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	var para models.OrderDetailsCondition

	para.Id = utility.StrToInt(r.URL.Query().Get("id"))               // take id for url
	para.CompanyId = utility.StrToInt(r.URL.Query().Get("CompanyId")) // take company_id for url

	parameters := models.OrderDetails{}.GetFilterOrderDetails(para)
	result, err := models.OrderDetails{}.OrderDetailsGet(parameters)
	if err != nil {
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		utility.RenderJsonResponse(w, r, response)
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
	utility.RenderJsonResponse(w, r, response)
	return false
}
