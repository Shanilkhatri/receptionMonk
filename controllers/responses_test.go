package controllers

import (
	"bytes"
	"database/sql"
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

// TEST #1 - PUT RESPONSE with correct Struct and also correct Data
func TestResponsePutWithCorrectData(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":           0,
		"response":     "This is how you solve a certain thing",
		"ticketId":     10,
		"responseTime": 1698661052,
		"type":         "",
		"respondeeId":  0,
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
		UserId:        6, // userId corressponding ticket is diff
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
	dbmock.ExpectExec("INSERT INTO `responses`").WillReturnResult(sqlmock.NewResult(1, 1))

	dbmock.ExpectExec("UPDATE tickets SET").WillReturnResult(sqlmock.NewResult(1, 1))

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
	userdetails.CompanyID = 5
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
	request := httptest.NewRequest(http.MethodPut, "/responses", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutResponse(w, request)
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
	if data.Message != "Response sent successfuly!" || w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 - putResponse with incorrect-data
// -> rightnow the only incorrect data would be ""responseTime"".
// -> expecting a 400
func TestResponsePutWithIncorrectDateString(t *testing.T) {
	jsonData := map[string]interface{}{
		"id":           0,
		"response":     "This is how you solve a certain thing",
		"ticketId":     10,
		"responseTime": "sddf", // incorrect data type
		"type":         "",
		"respondeeId":  0,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

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
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/responses", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutResponse(w, request)

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

// TEST #3 putResponse with invalid access rights
// -> accountType: user trying to put response to some other ticket
// -> expecting a 403
func TestResponsePutWithIncorrectAccessRights(t *testing.T) {
	jsonData := map[string]interface{}{
		"id":           0,
		"response":     "This is how you solve a certain thing",
		"ticketId":     10,
		"responseTime": 1698661052,
		"type":         "",
		"respondeeId":  0,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as user
	// well generally there would be all the details in this struct
	// but for conducting tests I have just populated the needful feilds
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user"
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

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	expectedTicketRow := models.Tickets{
		Id:            3,
		UserId:        6, // userId corressponding ticket is diff
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

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/responses", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutResponse(w, request)

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

	if w.Result().StatusCode != 403 || data.Message != "Unauthorized access! You are not allowed to make this request" {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Result().StatusCode)
	}
}

// TEST #4 putResponse with valid access rights
// -> accountType: user trying to put response to his own ticket
// -> expecting a 400
func TestResponsePutWithErrInSql(t *testing.T) {
	jsonData := map[string]interface{}{
		"id":           0,
		"response":     "This is how you solve a certain thing",
		"ticketId":     10,
		"responseTime": 1698661052,
		"type":         "",
		"respondeeId":  0,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as user
	// well generally there would be all the details in this struct
	// but for conducting tests I have just populated the needful feilds
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user"
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

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE id = ?").WillReturnError(errors.New("Some sql error"))

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/responses", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutResponse(w, request)

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

	if w.Result().StatusCode != 400 || data.Message != "Cannot send response at the moment! Please try again." {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Result().StatusCode)
	}
}

// GetResponse
// testcase #1 GetResponse with correct Struct and also correct Data,show 200.
func TestGetresponseWithCorrectStruct(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]string{
		"ticketId": "1",
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user" // type set to user
	userdetails.CompanyID = 1

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	expectedTicket := models.Tickets{
		Id:        1,
		UserId:    1,
		CompanyId: 1,
	}
	// make expected user a row that will be returned
	rows1 := sqlmock.NewRows([]string{"id", "userId", "companyId"}).AddRow(expectedTicket.Id, expectedTicket.UserId, expectedTicket.CompanyId)
	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE id = ?").WillReturnRows(rows1)
	// mocking the user detail I expect from the Db op
	expectedResponse := models.Response{
		Id:           1,
		Response:     "",
		TicketId:     1,
		ResponseTime: 1698747288,
		Type:         "",
		RespondeeId:  1, // This is basically userId
	}
	// make expected user a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "response", "ticketId", "ResponseTime", "Type", "RespondeeId"}).AddRow(expectedResponse.Id, expectedResponse.Response, expectedResponse.TicketId, expectedResponse.ResponseTime, expectedResponse.Type, expectedResponse.RespondeeId)

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* from responses ").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/response?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetResponse(w, request)

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

	if data.Message != "Returns all matching users." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #2 GetResponse ,if user id is different,show 403.
func TestGetresponse(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]string{
		"ticketId": "1",
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user" // type set to user
	userdetails.CompanyID = 1

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	expectedTicket := models.Tickets{
		Id:        1,
		UserId:    2,
		CompanyId: 1,
	}
	// make expected user a row that will be returned
	rows1 := sqlmock.NewRows([]string{"id", "userId", "companyId"}).AddRow(expectedTicket.Id, expectedTicket.UserId, expectedTicket.CompanyId)
	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE id = ?").WillReturnRows(rows1)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/response?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetResponse(w, request)

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

	if data.Message != "Unauthorized access! You are not allowed to make this request." && w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #3 GetResponse, if ticket id is empty string,show 400.
func TestGetResponseWithoutTicketId(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]string{
		"ticketId": "",
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user" // type set to user
	userdetails.CompanyID = 1

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}

	// here we will prepare the url with parameters to pass to our request
	url := "/response?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetResponse(w, request)

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

	if data.Message != "Bad request! No ticketId found." && w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #4 GetResponse, empty user Struct,show 403.
func TestGetResponseWithEmptyUser(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]string{
		"ticketId": "1",
	}

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool: false,
	}

	// here we will prepare the url with parameters to pass to our request
	url := "/response?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetResponse(w, request)

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

	if data.Message != "Unauthorized access! You are not allowed to make this request" && w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #5 GetResponse if ticket table return empty rows,show 400.
func TestGetresponseReturnEmptyRows(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]string{
		"ticketId": "1",
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user" // type set to user
	userdetails.CompanyID = 1

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// make expected user a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "userId", "companyId"})
	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE id = ?").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/response?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetResponse(w, request)

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

	if data.Message != "Cannot get ticket data at the moment! Please try again." && w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #6 GetResponse return sql errors,show 500.
func TestGetresponseWithSqlErrors(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]string{
		"ticketId": "1",
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user" // type set to user
	userdetails.CompanyID = 1

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	expectedTicket := models.Tickets{
		Id:        1,
		UserId:    1,
		CompanyId: 1,
	}
	// make expected user a row that will be returned
	rows1 := sqlmock.NewRows([]string{"id", "userId", "companyId"}).AddRow(expectedTicket.Id, expectedTicket.UserId, expectedTicket.CompanyId)
	dbmock.ExpectQuery("SELECT \\* FROM `tickets` WHERE id = ?").WillReturnRows(rows1)

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM `responses` WHERE id = ?").WillReturnError(sql.ErrNoRows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/response?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetResponse(w, request)

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

	if data.Message != "Internal server error, Any serious issues which cannot be recovered from." && w.Result().StatusCode != 500 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}
