package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reakgo/utility"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

// TEST #1 - PUT TICKET with correct Struct and also correct Data
func TestTicketPutWithCorrectData(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":            0,
		"userId":        0,
		"email":         "",
		"customerName":  "",
		"createdTime":   0,
		"lastUpdatedOn": 0,
		"status":        "open",
		"query":         "how to do a certain thing??",
		"feedback":      "",
		"lastResponse":  "",
		"companyId":     0,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("INSERT INTO `tickets`").WillReturnResult(sqlmock.NewResult(1, 1))

	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// mocking token payload with userDetails as owner
	// well generally there would be all the details in this struct
	// but for conducting tests I have just populated the needful feilds
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 3
	userdetails.Email = "xyz@ymail.com"
	userdetails.Name = "shan"
	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/tickets", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutTicket(w, request)
	err = dbmock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expectations were not met %s", err)
	}

	// Read the response body into a byte slice
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		// here i am just logging the error
		log.Println(err)
	}

	var data utility.AjaxResponce
	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(body, &data); err != nil {
		// Handle the JSON unmarshaling error
		log.Println("error", err)
	}
	log.Println("data.Message: ", data.Message)
	if data.Message != "Ticket raised successfuly!" || w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 - putTicket with incorrect-data
// -> rightnow the only incorrect data would be ""query"" as rest all of it would be set at controller
// -> expecting a 400
func TestTicketPutWithIncorrectDateString(t *testing.T) {
	jsonData := map[string]interface{}{
		"id":            0,
		"userId":        0,
		"email":         "",
		"customerName":  "",
		"createdTime":   0,
		"lastUpdatedOn": 0,
		"status":        "open",
		"query":         8786786, // incorrect data type
		"feedback":      "",
		"lastResponse":  "",
		"companyId":     0,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as owner
	// well generally there would be all the details in this struct
	// but for conducting tests I have just populated the needful feilds
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 3
	userdetails.Email = "xyz@ymail.com"
	userdetails.Name = "shan"
	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutTicket(w, request)

	// Read the response body into a byte slice
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		// here i am just logging the error
		log.Println(err)
	}

	var data utility.AjaxResponce
	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(body, &data); err != nil {
		// Handle the JSON unmarshaling error
		log.Println("error", err)
	}
	// log.Println("data.Message: ", data.Message)

	if w.Result().StatusCode != 400 || data.Message != "Please check all fields correctly and try again." {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Result().StatusCode)
	}
}
