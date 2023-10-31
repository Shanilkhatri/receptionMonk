package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
	"time"
)

func PutResponse(w http.ResponseWriter, r *http.Request) {
	// preparing an ajaxResponse
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var responseStruct models.Response
	err := utility.StrictParseDataFromJson(r, &responseStruct)
	log.Println("responseStruct: ", responseStruct)
	if err != nil {
		log.Println("Unable to decode json")
		// utility.Logger(err)
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// after parsing data we are now checking who is raising the ticket
	isok, userDetails := Utility.CheckTokenPayloadAndReturnUser(r)
	if !isok {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	// get ticket by Id
	ticketStruct, err := models.Tickets{}.GetTicketById(responseStruct.TicketId)
	if ticketStruct.Status == "closed" || ticketStruct.Status == "archive" {
		response.Status = "400"
		response.Message = "This ticket is no longer active."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	if userDetails.AccountType == "user" || userDetails.AccountType == "owner" {
		if err != nil {
			log.Println("error: ", err)
			response.Status = "400"
			response.Message = "Cannot send response at the moment! Please try again."
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
		// users and owners can only reply to their tickets
		if ticketStruct.UserId != userDetails.ID {
			response.Status = "403"
			response.Message = "Unauthorized access! You are not allowed to make this request"
			utility.RenderJsonResponse(w, r, response, 403)
			return
		}
	}
	// preparing response struct to insert into response(customer data from token + from form)
	responseStruct.ResponseTime = time.Now().Unix()
	responseStruct.RespondeeId = userDetails.ID
	// responseStruct.TicketId will come from frontEnd as json
	// responseStruct.Response will come from frontEnd as json
	switch userDetails.AccountType {
	case "user", "owner":
		responseStruct.Type = "customer"
	case "super-admin":
		responseStruct.Type = "super-admin"
	case "employee":
		responseStruct.Type = "employee"
	}

	// initialising a transaction
	tx := utility.Db.MustBegin()

	// to handle any panic
	defer func() {
		if recover := recover(); recover != nil {
			log.Println("panic occured: ", recover)
			tx.Rollback()
			response.Message = "An internal error occurred, please try again"
			utility.RenderJsonResponse(w, r, response, 500)
		}
	}()
	// calling the model for insertion
	isOk := models.Response{}.PutResponse(responseStruct, tx)
	if !isOk {
		response.Status = "400"
		response.Message = "Couldn't send response at the moment!"
		tx.Rollback()
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// changing status of ticket to in_process
	if ticketStruct.Status != "in_process" {
		// setting status to in_process
		ticketStruct.Status = "in_process"
		isok, _ := models.Tickets{}.PostTicket(ticketStruct, tx)
		if !isok {
			response.Status = "400"
			response.Message = "Couldn't change ticket status at the moment!"
			tx.Rollback()
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
	}
	tx.Commit()
	response.Status = "200"
	response.Message = "Response sent successfuly!"
	utility.RenderJsonResponse(w, r, response, 200)
}

func GetResponse(w http.ResponseWriter, r *http.Request) {
	// preparing an ajaxResponse
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	var responseCondition models.ResponseCondition
	isOk, userDetails := Utility.CheckTokenPayloadAndReturnUser(r)
	if !isOk {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	responseCondition.TicketId = int64(utility.StrToInt(r.URL.Query().Get("ticketId")))

	if responseCondition.TicketId == 0 {
		response.Status = "400"
		response.Message = "Bad request! No ticketId found."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// get ticket by Id
	ticketStruct, err := models.Tickets{}.GetTicketById(responseCondition.TicketId)
	if err != nil {
		response.Status = "400"
		response.Message = "Cannot get ticket data at the moment! Please try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	if userDetails.AccountType == "user" || userDetails.AccountType == "owner" {
		if err != nil {
			log.Println("error: ", err)
			response.Status = "400"
			response.Message = "Cannot get response at the moment! Please try again."
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
		// users and owners can only reply to their tickets
		if ticketStruct.UserId != userDetails.ID {
			response.Status = "403"
			response.Message = "Unauthorized access! You are not allowed to make this request."
			utility.RenderJsonResponse(w, r, response, 403)
			return
		}
	}
	result, err := models.Response{}.GetResponses(responseCondition.TicketId)
	if err != nil {
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		utility.RenderJsonResponse(w, r, response, 500)
		return
	} else if len(result) == 0 {
		response.Status = "400"
		response.Message = "No result were found for this search. Either record is not present or you are not authorized to access this users data"
		utility.RenderJsonResponse(w, r, response, 400)
		return
	} else {
		response.Status = "200"
		response.Message = "Returns all matching users."
		response.Payload = result // Set the user data in the response payload
	}
	utility.RenderJsonResponse(w, r, response, 200)
}
