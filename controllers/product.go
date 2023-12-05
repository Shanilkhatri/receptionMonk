package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
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

func IsValidProductStruct(productStruct models.Products) bool {
	if productStruct.Name != "" && productStruct.Description != "" && productStruct.Status != "" && productStruct.Price > 0 && productStruct.PlanValidity > 0 {
		return true
	}
	return false
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	var productStruct models.Products
	err := Helper.StrictParseDataFromJson(r, &productStruct)
	if err != nil {
		// Helper.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		return
	}
	isok, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isok {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	if userDetails.AccountType == "" {
		// Helper.Logger(err)
		response.Status = "403"
		response.Message = "You cannot add the product because you are not an admin."
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	if !IsValidProductStruct(productStruct) {
		// Helper.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Either required fields are empty or contain invalid data type"
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	ok := models.Products{}.PutProduct(productStruct)
	if !ok {
		response.Status = "400"
		response.Message = "Unable to add product at the moment! Please try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return

	}
	response.Status = "200"
	response.Message = "successfully added the product."
	Helper.RenderJsonResponse(w, r, response, 400)

}
func PostProduct(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	var productStruct models.Products
	err := Helper.StrictParseDataFromJson(r, &productStruct)
	if err != nil {
		// Helper.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		return
	}
	isok, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isok {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	if userDetails.AccountType == "" {
		// Helper.Logger(err)
		response.Status = "403"
		response.Message = "You cannot update the product because you are not an admin."
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	ok, err := models.Products{}.PostProduct(productStruct)
	if !ok || err != nil {
		response.Status = "400"
		response.Message = "Unable to add product at the moment! Please try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return

	}
	response.Status = "200"
	response.Message = "successfully update the product."
	Helper.RenderJsonResponse(w, r, response, 400)

}
