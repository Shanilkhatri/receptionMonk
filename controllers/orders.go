package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
	"time"

	"github.com/go-sql-driver/mysql"
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
			Helper.RenderJsonResponse(w, r, response, 500)
		}
	}()

	// var filter models.GetOrderFilter
	//  get params from query
	// get Prod Id from url
	isOk, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
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
				paramParsed, err := Helper.StrToInt64(paramValue)
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
			Helper.RenderJsonResponse(w, r, response, 403)
			return
		}
		// when user tries to access some other account
		if userDetails.AccountType == "user" {
			if paramMap["userId"] != 0 && userDetails.ID != int64(paramMap["userId"]) {
				log.Println("under user not authed")
				response.Status = "403"
				response.Message = "You are not authorized for this request"
				Helper.RenderJsonResponse(w, r, response, 403)
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
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		} else {
			response.Status = "200"
			response.Message = "Results found successfully"
			response.Payload = result
			Helper.RenderJsonResponse(w, r, response, 200)
			return
		}
	}
	response.Message = "Failed to authorize at the moment. Please Login again and try!"
	Helper.RenderJsonResponse(w, r, response, 500)
}

func PostOrder(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	var orderStruct models.Orders
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

	err := Helper.StrictParseDataFromJson(r, &orderStruct)
	if err != nil {
		log.Println(err)
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}
	//check validation.
	boolType := OrderValidationCheck(orderStruct)
	if boolType {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	tx := utility.Db.MustBegin()
	// get planValidity.
	planValidity, err := models.Orders{}.GetSingleProduct(orderStruct.ProductId, tx)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}
	// calculate Expiry.
	orderStruct.Expiry = ExpiryCalculate(orderStruct.PlacedOn, planValidity)
	boolValue, err := models.Orders{}.PostOrder(orderStruct, tx)

	if !boolValue || err != nil {
		log.Println(err)
		tx.Rollback()
		response.Status = "500"
		response.Message = "Internal server error1, Any serious issues which cannot be recovered from."
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
		response.Message = "Order Update successfully."
	}

	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}

func ExpiryCalculate(placedOn int64, planValidityInDays int64) int64 {

	planValiditySeconds := int64(3600 * 24 * planValidityInDays)

	validityDuration := time.Duration(planValiditySeconds) * time.Second

	placedOnTime := time.Unix(placedOn, 0)

	expiryTime := placedOnTime.Add(validityDuration)

	// Convert expiryTime back to Unix timestamp
	expiry := expiryTime.Unix()

	return expiry

}

func OrderValidationCheck(orderStruct models.Orders) bool {
	switch {
	case orderStruct.ProductId <= 0:
		return true
	case orderStruct.PlacedOn <= 0:
		return true
	case orderStruct.Price <= 0:
		return true
	case orderStruct.Buyer <= 0:
		return true
	case orderStruct.Status == "":
		return true
	default:
		return false
	}
}

func PutOrder(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	var orderStruct models.Orders
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

	err := Helper.StrictParseDataFromJson(r, &orderStruct)
	if err != nil {
		log.Println(err)
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	//check validation.
	boolType := OrderValidationCheck(orderStruct)
	if boolType {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}
	tx := utility.Db.MustBegin()
	//get planValidity. Db.leagues.
	planValidity, err := models.Orders{}.GetSingleProduct(orderStruct.ProductId, tx)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}
	//calculate Expiry.
	orderStruct.Expiry = ExpiryCalculate(orderStruct.PlacedOn, planValidity)

	//put data in table.
	boolValue, err := models.Orders{}.PutOrder(orderStruct, tx)
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
		response.Message = "Order added successfully."
	}

	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}

func OrderDelete(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}

	orderId := Helper.StrToInt(r.URL.Query().Get("id"))

	if orderId <= 0 {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	boolType, err := models.Orders{}.OrderDelete(orderId)
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
	response.Message = "Order deleted successfully."
	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}
