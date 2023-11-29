package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
	"time"
)

func PutTicket(w http.ResponseWriter, r *http.Request) {
	// preparing an ajaxResponse
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var ticketStruct models.Tickets
	err := Helper.StrictParseDataFromJson(r, &ticketStruct)
	log.Println("ticketStruct: ", ticketStruct)
	if err != nil {
		log.Println("Unable to decode json")
		// utility.Logger(err)
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	// after parsing data we are now checking who is raising the ticket
	isok, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isok {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	// preparing ticket struct to insert into Tickets(customer data from token + from form)
	ticketStruct.UserId = userDetails.ID
	ticketStruct.CustomerName = userDetails.Name
	ticketStruct.Email = userDetails.Email
	ticketStruct.CompanyId = userDetails.CompanyID
	// setting feedback = no_feedback
	ticketStruct.FeedBack = "no_feedback"
	// setting createdTime & lastUpdatedOn to current time
	ticketStruct.CreatedTime = time.Now().Unix()
	ticketStruct.LastUpdatedOn = time.Now().Unix()
	// setting status to "open" initially
	ticketStruct.Status = "open"
	// setting last response to empty initially
	ticketStruct.LastResponse = ""

	// initialising a transaction
	tx := utility.Db.MustBegin()

	// to handle any panic
	defer func() {
		if recover := recover(); recover != nil {
			log.Println("panic occured: ", recover)
			tx.Rollback()
			response.Message = "An internal error occurred, please try again"
			Helper.RenderJsonResponse(w, r, response, 500)
		}
	}()
	// calling the model for insertion
	isOk := models.Tickets{}.PutTicket(ticketStruct, tx)
	if !isOk {
		response.Status = "400"
		response.Message = "Couldn't create a ticket at the moment!"
		tx.Rollback()
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	tx.Commit()
	response.Status = "200"
	response.Message = "Ticket raised successfuly!"
	Helper.RenderJsonResponse(w, r, response, 200)
}

func PostTicket(w http.ResponseWriter, r *http.Request) {
	// preparing an ajaxResponse
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var ticketStruct models.Tickets
	err := Helper.StrictParseDataFromJson(r, &ticketStruct)
	log.Println("ticketStruct: ", ticketStruct)
	if err != nil {
		log.Println("Unable to decode json")
		Helper.Logger(err)
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	// validation only employee & super-admin can update ticket
	isok, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isok || userDetails.AccountType != "employee" && userDetails.AccountType != "super-admin" {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request."
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	//  will only open this condition when :
	// -> only super-admin can edit/update any ticket and,
	// -> employee can only update tickets to the company he belongs
	// -> for now both can update any ticket
	// if userDetails.AccountType != "super-admin" && userDetails.CompanyID != ticketStruct.CompanyId {
	// 	response.Status = "403"
	// 	response.Message = "Unauthorized access! You are not allowed to make this request"
	// 	utility.RenderJsonResponse(w, r, response, 403)
	// 	return
	// }

	// setting lastUpdated on to current timestamp
	ticketStruct.LastUpdatedOn = time.Now().Unix()

	// get the existing ticket
	ticketData, err := models.Tickets{}.GetTicketById(ticketStruct.Id)
	if err != nil {
		errStr := Helper.GetSqlErrorString(err)
		log.Println("sqlError: ", errStr)
		response.Status = "400"
		response.Message = errStr
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	log.Println("ticketData: ", ticketData)
	// copying all the unchanged fields to ticketStruct
	if Helper.FillEmptyFieldsForPostUpdate(ticketData, &ticketStruct) {
		tx := utility.Db.MustBegin()
		log.Println("ticketStruct after flip: ", ticketStruct)
		// to handle any panic
		defer func() {
			if recover := recover(); recover != nil {
				log.Println("panic occured: ", recover)
				tx.Rollback()
				response.Message = "An internal error occurred, please try again"
				Helper.RenderJsonResponse(w, r, response, 500)
			}
		}()
		// now begin update by id next
		_, err := models.Tickets{}.PostTicket(ticketStruct, tx)
		if err != nil {
			log.Println("err when updating ticket: ", err)
			response.Status = "400"
			response.Message = "Unable to update records at the moment! Please try again."
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		}
		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			response.Status = "400"
			response.Message = "Unable to update records at the moment! Please try again."
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		}
		response.Status = "200"
		response.Message = "Record updated successfully."
		Helper.RenderJsonResponse(w, r, response, 200)
		return
	}
}

func GetTicket(w http.ResponseWriter, r *http.Request) {
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

	//  get params from query
	isOk, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if isOk {
		queryParams := r.URL.Query()

		paramMap := map[string]int64{
			"id":            0,
			"userId":        0,
			"companyId":     0,
			"createdTime":   0,
			"lastUpdatedOn": 0,
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
		// validation (anyone other than employee, user, owner & super-admin will not be allowed)
		if userDetails.AccountType != "employee" && userDetails.AccountType != "user" && userDetails.AccountType != "owner" && userDetails.AccountType != "super-admin" {
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

		param := models.TicketsCondition{
			Tickets: models.Tickets{
				Id:            paramMap["id"],
				UserId:        paramMap["userId"],
				CompanyId:     paramMap["companyId"],
				CreatedTime:   paramMap["createdTime"],
				LastUpdatedOn: paramMap["lastUpdatedOn"],
			},
		}

		params := models.Tickets{}.GetParamsForFilterTicketsData(param)
		result, err := models.Orders{}.GetTickets(params, tx)
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

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}

	ticketId := Helper.StrToInt(r.URL.Query().Get("id"))
	if ticketId <= 0 {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	isok, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isok {
		response.Status = "403"
		response.Message = "You are not authorized to make this request."
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	if userDetails.AccountType == "user" {
		// get the ticket he wants to delete and check if the userId matches
		ticketData, err := models.Tickets{}.GetTicketById(int64(ticketId))
		if err != nil {
			response.Status = "400"
			response.Message = Helper.GetSqlErrorString(err)
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		}
		if ticketData.UserId != userDetails.ID {
			response.Status = "403"
			response.Message = "You are not authorized to make this request."
			Helper.RenderJsonResponse(w, r, response, 403)
			return
		}
	}
	if userDetails.AccountType == "owner" {
		// get the ticket he wants to delete and check if the userId matches
		ticketData, err := models.Tickets{}.GetTicketById(int64(ticketId))
		if err != nil {
			response.Status = "400"
			response.Message = Helper.GetSqlErrorString(err)
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		}
		if ticketData.CompanyId != userDetails.CompanyID {
			response.Status = "403"
			response.Message = "You are not authorized to make this request."
			Helper.RenderJsonResponse(w, r, response, 403)
			return
		}
	}
	_, err := models.Tickets{}.DeleteTicket(ticketId)
	if err != nil {
		response.Status = "400"
		response.Message = Helper.GetSqlErrorString(err)
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	response.Status = "200"
	response.Message = "Ticket deleted successfully."
	Helper.RenderJsonResponse(w, r, response, 200)
}
