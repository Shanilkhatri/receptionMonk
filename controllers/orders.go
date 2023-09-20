package controllers

import (
	"encoding/json"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
)

func OrdersJsonDecoder(r *http.Request) (models.Orders, error) {
	var orderStruct models.Orders
	err := json.NewDecoder(r.Body).Decode(&orderStruct)
	return orderStruct, err
}

func GetOrders(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}

	return false
}
