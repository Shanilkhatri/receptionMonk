package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	tx := utility.Db.MustBegin()
	// for handling any panic
	defer func() {
		if recover := recover(); recover != nil {
			log.Println("panic occured: ", recover)
			tx.Rollback()
			response.Message = "An internal error occurred, please try again"
			utility.RenderJsonResponse(w, r, response, 500)
		}
	}()

	// var filter models.GetOrderFilter
	//  get params from query
	// get Prod Id from url
	isOk, userDetails := utility.CheckTokenPayloadAndReturnUser(r)
	log.Println("userDetails: ", userDetails)
	if isOk {
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

		// validation (anyone other than user, owner & super-admin will not be allowed)
		if userDetails.AccountType != "user" && userDetails.AccountType != "owner" && userDetails.AccountType != "super-admin" {
			response.Status = "403"
			response.Message = "You are not authorized for this request"
			utility.RenderJsonResponse(w, r, response, 403)
			return
		}
		// when user tries to access some other account
		if userDetails.AccountType == "user" {
			if paramMap["userId"] != 0 && userDetails.ID != int64(paramMap["userId"]) {
				log.Println("under user not authed")
				response.Status = "403"
				response.Message = "You are not authorized for this request"
				utility.RenderJsonResponse(w, r, response, 403)
				return
			} else {
				// id given from token data
				paramMap["userId"] = int64(userDetails.ID)
			}
		}
		if userDetails.AccountType == "owner" {
			if paramMap["userId"] != 0 || paramMap["id"] != 0 {
				paramMap["companyId"] = int64(userDetails.CompanyID)
			} else {
				// id given from token data
				paramMap["userId"] = int64(userDetails.ID)
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

		params := models.Orders{}.GetParamsForFilterOrderData(param)
		result, err := models.Orders{}.GetOrders(params, tx)
		if err != nil {
			log.Println(err)
		} else if len(result) == 0 {
			response.Status = "400"
			response.Message = "Either You don't have access or there isn't any record present! Please try again with valid parameters."
			utility.RenderJsonResponse(w, r, response, 400)
			return
		} else {
			response.Status = "200"
			response.Message = "Results found successfully"
			response.Payload = result
			utility.RenderJsonResponse(w, r, response, 200)
			return
		}
	}
	response.Message = "Failed to authorize at the moment. Please Login again and try!"
	utility.RenderJsonResponse(w, r, response, 500)
}
