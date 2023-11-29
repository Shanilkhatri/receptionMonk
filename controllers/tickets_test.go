package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reakgo/models"
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
	Helper = MockHelper{
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
	Helper = MockHelper{
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

// ---------------- TEST TICKET POST ------------------

// TEST #1 PostTicket with correct data
// -> expecting a 200

func TestTicketPostWithCorrectData(t *testing.T) {
	jsonData := map[string]interface{}{
		"id":           3,
		"userId":       1,
		"email":        "shasha@hmail.com",
		"customerName": "shani",
		// removed created time as it will always be the same
		"lastUpdatedOn": 1698300819,
		"status":        "open",
		"query":         "how to do a certain thing a certain way?",
		"feedback":      "no_feedback",
		"lastResponse":  "",
		"companyId":     2,
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

	expectedTicketRow := models.Tickets{
		Id:            3,
		UserId:        1,
		Email:         "shasha@hmail.com",
		CustomerName:  "shanu",
		CreatedTime:   1698214419,
		LastUpdatedOn: 1698214419,
		Status:        "open",
		Query:         "how to do a certain thing",
		FeedBack:      "no_feedback",
		LastResponse:  "",
		CompanyId:     2,
	}

	// make expected Ticket a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "userId", "email", "customerName", "createdTime", "lastUpdatedOn", "status", "query", "feedback", "lastResponse", "companyId"}).AddRow(expectedTicketRow.Id, expectedTicketRow.UserId, expectedTicketRow.Email, expectedTicketRow.CustomerName, expectedTicketRow.CreatedTime, expectedTicketRow.LastUpdatedOn, expectedTicketRow.Status, expectedTicketRow.Query, expectedTicketRow.FeedBack, expectedTicketRow.LastResponse, expectedTicketRow.CompanyId)

	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE id = ?").WillReturnRows(rows)
	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("UPDATE tickets SET userId=").WillReturnResult(sqlmock.NewResult(1, 1))

	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// mocking token payload with userDetails as employee
	// well generally there would be all the details in this struct
	// but for conducting tests I have just populated the needful feilds
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "employee"
	userdetails.CompanyID = 3
	userdetails.Email = "xyz@ymail.com"
	userdetails.Name = "shan"
	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
		MockGetSqlErrorString:                     "",
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/tickets", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostTicket(w, request)
	// check expectations
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

	if w.Result().StatusCode != 200 || data.Message != "Record updated successfully." {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 TICKET POST with incorrect data
// -> any accountType other than employee or super-admin cannot update tickets
// -> we'll mock the userdetails to give accountType "owner"
// -> expecting a 403
func TestTicketPostWithInCorrectData(t *testing.T) {
	jsonData := map[string]interface{}{
		"id":           3,
		"userId":       1,
		"email":        "shasha@hmail.com",
		"customerName": "shani",
		// removed created time as it will always be the same
		"lastUpdatedOn": 1698300819,
		"status":        "open",
		"query":         "how to do a certain thing a certain way?",
		"feedback":      "no_feedback",
		"lastResponse":  "",
		"companyId":     2,
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
	userdetails.AccountType = "owner" // unauthorised account type owner
	userdetails.CompanyID = 3
	userdetails.Email = "xyz@ymail.com"
	userdetails.Name = "shan"
	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/tickets", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostTicket(w, request)

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

	if w.Result().StatusCode != 403 || data.Message != "Unauthorized access! You are not allowed to make this request." {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Result().StatusCode)
	}
}

// TEST #3 PostTicket with correct data
// -> this time our sql op will throw an error
// -> expecting a 400

func TestTicketPostWithCorrectDataButErrAtSQL(t *testing.T) {
	jsonData := map[string]interface{}{
		"id":           3,
		"userId":       1,
		"email":        "shasha@hmail.com",
		"customerName": "shani",
		// removed created time as it will always be the same
		"lastUpdatedOn": 1698300819,
		"status":        "open",
		"query":         "how to do a certain thing a certain way?",
		"feedback":      "no_feedback",
		"lastResponse":  "",
		"companyId":     2,
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

	expectedTicketRow := models.Tickets{
		Id:            3,
		UserId:        1,
		Email:         "shasha@hmail.com",
		CustomerName:  "shanu",
		CreatedTime:   1698214419,
		LastUpdatedOn: 1698214419,
		Status:        "open",
		Query:         "how to do a certain thing",
		FeedBack:      "no_feedback",
		LastResponse:  "",
		CompanyId:     2,
	}

	// make expected Ticket a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "userId", "email", "customerName", "createdTime", "lastUpdatedOn", "status", "query", "feedback", "lastResponse", "companyId"}).AddRow(expectedTicketRow.Id, expectedTicketRow.UserId, expectedTicketRow.Email, expectedTicketRow.CustomerName, expectedTicketRow.CreatedTime, expectedTicketRow.LastUpdatedOn, expectedTicketRow.Status, expectedTicketRow.Query, expectedTicketRow.FeedBack, expectedTicketRow.LastResponse, expectedTicketRow.CompanyId)

	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE id = ?").WillReturnRows(rows)
	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("UPDATE tickets SET userId=").WillReturnError(errors.New("Some sql error"))

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// mocking token payload with userDetails as employee
	// well generally there would be all the details in this struct
	// but for conducting tests I have just populated the needful feilds
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "employee"
	userdetails.CompanyID = 3
	userdetails.Email = "xyz@ymail.com"
	userdetails.Name = "shan"
	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/tickets", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostTicket(w, request)
	// check expectations
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

	if w.Result().StatusCode != 400 || data.Message != "Unable to update records at the moment! Please try again." {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// -------------------------- TEST TICKET GET -------------------------------

// TEST #1 - GetTickets with correct parameters
// -> There are four access levels based on "accountType" to get orders info and they are:
// -> when AccountType : user -> this account can only get Tickets related to his userId
// -> when AccountType : owner -> this account can get tickets related to his userId as well as all users working in his company
// -> when AccountType : employee -> this account can get orders info related to anyone(highest level of access rights)
// -> when AccountType : super-admin -> this account can get orders info related to anyone(highest level of access rights)
// -> for the first test we'll mock the user details from the token and make accountType : user
// -> we'll also mock the Db to perform ops of fetch
// expecting a 200
func TestGetTicketWithTypeUser(t *testing.T) {
	// parameters to be send on the basis of which we can filter tickets and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]string{
		"id":            "",
		"createdTime":   "",
		"lastUpdatedOn": "",
		"companyId":     "",
		"userId":        "3", // getting data for userId #1
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "employee" // type set to user
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// mocking the expected Ticket
	expectedTicketRow := models.Tickets{
		Id:            1,
		UserId:        3,
		Email:         "shasha@hmail.com",
		CustomerName:  "shanu",
		CreatedTime:   1698214419,
		LastUpdatedOn: 1698214419,
		Status:        "open",
		Query:         "how to do a certain thing",
		FeedBack:      "no_feedback",
		LastResponse:  "",
		CompanyId:     2,
	}
	// make expected Ticket a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "userId", "email", "customerName", "createdTime", "lastUpdatedOn", "status", "query", "feedback", "lastResponse", "companyId"}).AddRow(expectedTicketRow.Id, expectedTicketRow.UserId, expectedTicketRow.Email, expectedTicketRow.CustomerName, expectedTicketRow.CreatedTime, expectedTicketRow.LastUpdatedOn, expectedTicketRow.Status, expectedTicketRow.Query, expectedTicketRow.FeedBack, expectedTicketRow.LastResponse, expectedTicketRow.CompanyId)

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE 1=1").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/tickets?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetTicket(w, request)
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

	if data.Message != "Results found successfully" && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 - GetTickets with correct parameters
// -> user tries to access someone elses tickets
// -> for the second test we'll mock the user details from the token and make accountType : user
// -> we won't mock DB as we expect to get error unauhtorized access before the DB ops
// expecting a 403
func TestGetTicketsWithTypeUserAccessingAnotherUser(t *testing.T) {
	// parameters to be send on the basis of which we can filter tickets and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]string{
		"id":        "",
		"date_from": "",
		"date_to":   "",
		"companyId": "",
		"userId":    "2", // getting data for userId #2
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1               // userId #1 trying to access records of userId #2
	userdetails.AccountType = "user" // type set to user
	userdetails.CompanyID = 1

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// // open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/orders?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetTicket(w, request)

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

	if data.Message != "You are not authorized for this request" && w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Result().StatusCode)
	}
}

// TEST #3 - GetTickets with correct parameters
// -> owner tries to access someone else's tickets of diff company
// -> for the third test we'll mock the user details from the token and make accountType : owner
// -> we will mock DB as we expect to get result(length)=0
// expecting a 400
func TestGetTicketsWithTypeOwnerAccessingAnotherUserOfDiffCompany(t *testing.T) {
	// parameters to be send on the basis of which we can filter tickets and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]string{
		"id":        "",
		"date_from": "",
		"date_to":   "",
		"companyId": "",
		"userId":    "2", // getting data for userId #2
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1                // userId #1 trying to access records of userId #2
	userdetails.AccountType = "owner" // type set to user
	userdetails.CompanyID = 1

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// // open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// make tickets row that will be returned but empty
	rows := sqlmock.NewRows([]string{"id", "userId", "email", "customerName", "createdTime", "lastUpdatedOn", "status", "query", "feedback", "lastResponse", "companyId"})

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE 1=1").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/orders?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetTicket(w, request)

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

	if data.Message != "Either You don't have access or there isn't any record present! Please try again with valid parameters." || w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Result().StatusCode)
	}
}

// // TEST #4 - GetTickets with correct parameters
// // -> owner tries to access someone else's tickets of same company
// // -> for the fourth test we'll mock the user details from the token and make accountType : owner
// // -> we will mock DB as we expect to get result
// // expecting a 200
func TestGetTicketsWithTypeOwnerAccessingAnotherUserOfSameCompany(t *testing.T) {
	// parameters to be send on the basis of which we can filter tickets and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]string{
		"id":        "",
		"date_from": "",
		"date_to":   "",
		"companyId": "",
		"userId":    "3", // getting data for userId #3
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1                // userId #1 trying to access records of userId #2
	userdetails.AccountType = "owner" // type set to user
	userdetails.CompanyID = 1

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// // open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// mocking the expected Ticket
	expectedTicketRow := models.Tickets{
		Id:            1,
		UserId:        3,
		Email:         "shasha@hmail.com",
		CustomerName:  "shanu",
		CreatedTime:   1698214419,
		LastUpdatedOn: 1698214419,
		Status:        "open",
		Query:         "how to do a certain thing",
		FeedBack:      "no_feedback",
		LastResponse:  "",
		CompanyId:     1,
	}
	// make expected Ticket a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "userId", "email", "customerName", "createdTime", "lastUpdatedOn", "status", "query", "feedback", "lastResponse", "companyId"}).AddRow(expectedTicketRow.Id, expectedTicketRow.UserId, expectedTicketRow.Email, expectedTicketRow.CustomerName, expectedTicketRow.CreatedTime, expectedTicketRow.LastUpdatedOn, expectedTicketRow.Status, expectedTicketRow.Query, expectedTicketRow.FeedBack, expectedTicketRow.LastResponse, expectedTicketRow.CompanyId)

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE 1=1").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/tickets?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetTicket(w, request)

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
	log.Println("msg: ", data.Message)
	if data.Message != "Results found successfully" || w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// ---------------------TESTs FOR DELETE TICKET----------------------

// -> TicketId should be sent in query
// -> AccountType: user have access to this route, but he/she can only delete there tickets.
// -> AccountType: owner can delete tickets under his organization and his own tickets.
// -> AccountType: super-admin can delete any Ticket.

// TEST #1 DeleteTicket with correct access rights.
// AccountType: owner trying to delete a user's Ticket that belongs to his own company
// expecting a 200
func TestTicketDeleteWithOwnerDelUserOfSameComp(t *testing.T) {
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 5
	userdetails.AccountType = "owner" // type set to owner
	userdetails.CompanyID = 2         // companyId set same as user's

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println("mockErr: ", err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// mocking the expected Ticket
	expectedTicketRow := models.Tickets{
		Id:            1,
		UserId:        3,
		Email:         "shasha@hmail.com",
		CustomerName:  "shanu",
		CreatedTime:   1698214419,
		LastUpdatedOn: 1698214419,
		Status:        "open",
		Query:         "how to do a certain thing",
		FeedBack:      "no_feedback",
		LastResponse:  "",
		CompanyId:     2, // companyId same as owners
	}
	// make expected Ticket a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "userId", "email", "customerName", "createdTime", "lastUpdatedOn", "status", "query", "feedback", "lastResponse", "companyId"}).AddRow(expectedTicketRow.Id, expectedTicketRow.UserId, expectedTicketRow.Email, expectedTicketRow.CustomerName, expectedTicketRow.CreatedTime, expectedTicketRow.LastUpdatedOn, expectedTicketRow.Status, expectedTicketRow.Query, expectedTicketRow.FeedBack, expectedTicketRow.LastResponse, expectedTicketRow.CompanyId)

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE").WillReturnRows(rows)
	// I  expect a Delete Query execution and for that :
	dbmock.ExpectExec("UPDATE `tickets`").WillReturnResult(sqlmock.NewResult(1, 1))

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set TicketId to be deleted to 1
	ticketId := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/users?" + "id=" + ticketId

	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteTicket(w, request)

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
	if data.Message != "Ticket deleted successfully." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 DeleteTicket with incorrect access rights.
// AccountType: owner trying to delete a user's Ticket that belongs to other company
// expecting a 403
func TestTicketDeleteWithOwnerDelUserOfDiffComp(t *testing.T) {
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 5
	userdetails.AccountType = "owner" // type set to owner
	userdetails.CompanyID = 2         // companyId diff from user

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println("mockErr: ", err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// mocking the expected Ticket
	expectedTicketRow := models.Tickets{
		Id:            1,
		UserId:        3,
		Email:         "shasha@hmail.com",
		CustomerName:  "shanu",
		CreatedTime:   1698214419,
		LastUpdatedOn: 1698214419,
		Status:        "open",
		Query:         "how to do a certain thing",
		FeedBack:      "no_feedback",
		LastResponse:  "",
		CompanyId:     3, // companyId Diff from owners
	}
	// make expected Ticket a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "userId", "email", "customerName", "createdTime", "lastUpdatedOn", "status", "query", "feedback", "lastResponse", "companyId"}).AddRow(expectedTicketRow.Id, expectedTicketRow.UserId, expectedTicketRow.Email, expectedTicketRow.CustomerName, expectedTicketRow.CreatedTime, expectedTicketRow.LastUpdatedOn, expectedTicketRow.Status, expectedTicketRow.Query, expectedTicketRow.FeedBack, expectedTicketRow.LastResponse, expectedTicketRow.CompanyId)

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE").WillReturnRows(rows)
	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set userId to be deleted to 1
	ticketToDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/tickets?" + "id=" + ticketToDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteTicket(w, request)

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
	if data.Message != "You are not authorized to make this request." && w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #5 DeleteTicket with correct access rights.
// AccountType: super-admin trying to delete owners Ticket, and he will succeed, it's SUPER-ADMIN common!
// expecting a 200
func TestTicketDeleteWithOSuperAdmin(t *testing.T) {
	// mocking token payload with userDetails as super-admin
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 5
	userdetails.AccountType = "super-admin" // type set to super-admin
	userdetails.CompanyID = 2
	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println("mockErr: ", err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I  expect a Delete Query execution and for that :
	dbmock.ExpectExec("UPDATE `tickets`").WillReturnResult(sqlmock.NewResult(1, 1))

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set userId to be deleted to 1
	ticketToDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/tickets?" + "id=" + ticketToDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteTicket(w, request)

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
	if data.Message != "Ticket deleted successfully." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}
