package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
)

func OrdersJsonDecoder(r *http.Request) (models.Orders, error) {
	var orderStruct models.Orders
	err := json.NewDecoder(r.Body).Decode(&orderStruct)
	return orderStruct, err
}

func GetOrders(w http.ResponseWriter, r *http.Request) utility.AjaxResponce {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	tx := utility.Db.MustBegin()
	// for handling any panic
	defer func() {
		if recover := recover(); recover != nil {
			log.Println("panic occured: ", recover)
			tx.Rollback()
			response.Message = "An internal error occurred, please try again"
			utility.RenderTemplate(w, r, "", response)
		}
	}()
	// var filter models.GetOrderFilter
	//  get params from query
	// get Prod Id from url
	queryParams := r.URL.Query()

	paramMap := map[string]int64{
		"id":        0,
		"date_from": 0,
		"date_to":   0,
		"companyId": 0,
		"userId":    0,
	}
	for paramName := range paramMap {
		paramValue := queryParams.Get(paramName)
		if paramValue != "" {
			paramParsed, err := utility.StrToInt64(paramValue)
			if err != nil {
				log.Println(err)
			} else {
				paramMap[paramName] = paramParsed
			}
		}
	}
	param := models.OrderDataCondition{
		Orders: models.Orders{
			Id:        paramMap["id"],
			DateFrom:  paramMap["date_from"],
			DateTo:    paramMap["date_to"],
			UserId:    paramMap["userId"],
			CompanyId: paramMap["companyId"],
		},
	}
	params := Db.orders.GetParamsForFilterOrderData(param)
	result, err := Db.orders.GetOrders(params, tx)
	if err != nil {
		log.Println(err)
	} else if len(result) == 0 {
		response.Status = "success"
		response.Message = "No data were found for this search"
	} else {
		response.Status = "200"
		response.Message = "Results found successfully"
		response.Payload = result
		utility.RenderTemplate(w, r, "", response)
		tx.Commit()
		return response
	}
	tx.Rollback()
	utility.RenderTemplate(w, r, "", response)
	return response
}
