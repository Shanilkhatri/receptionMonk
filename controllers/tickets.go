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
	err := utility.StrictParseDataFromJson(r, &ticketStruct)
	log.Println("ticketStruct: ", ticketStruct)
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
			utility.RenderJsonResponse(w, r, response, 500)
		}
	}()
	// calling the model for insertion
	isOk := models.Tickets{}.PutTicket(ticketStruct, tx)
	if !isOk {
		response.Status = "400"
		response.Message = "Couldn't create a ticket at the moment!"
		tx.Rollback()
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	tx.Commit()
	response.Status = "200"
	response.Message = "Ticket raised successfuly!"
	utility.RenderJsonResponse(w, r, response, 200)
}

func PostTicket(w http.ResponseWriter, r *http.Request) {
	// preparing an ajaxResponse
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var ticketStruct models.Tickets
	err := utility.StrictParseDataFromJson(r, &ticketStruct)
	log.Println("ticketStruct: ", ticketStruct)
	if err != nil {
		log.Println("Unable to decode json")
		utility.Logger(err)
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// validation only employee & super-admin can update ticket
	isok, userDetails := Utility.CheckTokenPayloadAndReturnUser(r)
	if !isok || userDetails.ID != ticketStruct.Id {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	if userDetails.AccountType != "super-admin" && userDetails.CompanyID != ticketStruct.CompanyId {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	// get the existing ticket
	ticketData, err := models.Tickets{}.GetTicketById(ticketStruct.Id)
	log.Println("ticketData: ", ticketData)
	// copying all the unchanged fields to ticketStruct
	if utility.FillEmptyFieldsForPostUpdate(ticketData, &ticketStruct) {
		// now begin update by id next
		_, err := models.Tickets{}.PostTicket(ticketStruct)
		if err != nil {
			log.Println("err when updating ticket: ", err)
			response.Status = "400"
			response.Message = "Unable to update records at the moment! Please try again."
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
		response.Status = "200"
		response.Message = "Record updated successfully."
		utility.RenderJsonResponse(w, r, response, 200)
		return
	}

}
