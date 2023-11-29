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

// TEST #1 - PutKyc with correct parameters
// -> user tries to put someone elses kyc
// -> for the second test we'll mock the user details from the token and make accountType : user
// -> we won't mock DB as we expect to get error unauhtorized access before the DB ops
// expecting a 403
func TestPutKycDetailsWithTypeUserAccessingAnotherUser(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]interface{}{
		"doc_name":     "docName",
		"doc_pic_name": "img",
		"userid":       1,
		"companyId":    1,
	}
	requestBody, err := json.Marshal(queryParams)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 2                // userId #1 trying to access records of userId #2
	userdetails.AccountType = "owner" // type set to user
	userdetails.CompanyID = 2

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
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/responses", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	// Call your function with the mocks
	PutKycDetails(w, request)

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

// TEST #2 - PutKyc with incorrect parameters
// -> user tries to put incorrect kyc data(i had given the doc_name as a int here so you can change its type and check)
// -> we won't mock DB as we expect to get error unauhtorized access before the DB ops
// expecting a 400
func TestPutKycDetailsWithIncorrectData(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]interface{}{
		"doc_name":     22,
		"doc_pic_name": "img",
		"userid":       1,
		"companyId":    1,
	}
	requestBody, err := json.Marshal(queryParams)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 2                // userId #1 trying to access records of userId #2
	userdetails.AccountType = "owner" // type set to user
	userdetails.CompanyID = 2

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
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/responses", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	// Call your function with the mocks
	PutKycDetails(w, request)

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

	if data.Message != "Please check all fields correctly and try again." && w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Result().StatusCode)
	}
}

// TEST #3 - PUT kyc with correct Struct and also correct Data
// -> owner tries to add someone else's kyc of same company
// -> for the third test we'll mock the user details from the token and make accountType : owner
// -> we will mock DB as we expect to get result
// expecting a 200
func TestPutKycDetailsWithCorrectData(t *testing.T) {

	jsonData := map[string]interface{}{
		"userid":       2,
		"doc_name":     "aadhar_card",
		"doc_pic_name": "something.png",
		"companyId":    1,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("INSERT INTO `kyc_details`").WillReturnResult(sqlmock.NewResult(1, 1))

	dbmock.ExpectExec("UPDATE `authentication` SET").WillReturnResult(sqlmock.NewResult(1, 1))

	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// mocking token payload with userDetails as employee
	// well generally there would be all the details in this struct
	// but for conducting tests I have just populated the needful feilds
	var userdetails utility.UserDetails
	userdetails.ID = 2
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 1
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
	request := httptest.NewRequest(http.MethodPut, "/kyc", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutKycDetails(w, request)
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
	if data.Message != "Document upload successfully" || w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #4 - POST kyc with correct Struct and also correct Data
// -> owner tries to update someone else's kyc of same company
// -> for the forth test we'll mock the user details from the token and make accountType : owner
// -> we will mock DB as we expect to get result
// expecting a 200
func TestPostKycDetailsWithCorrectData(t *testing.T) {

	jsonData := map[string]interface{}{
		"userid":       2,
		"doc_name":     "pan_card",
		"doc_pic_name": "pan_something.png",
		"companyId":    1,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("UPDATE `kyc_details` SET").WillReturnResult(sqlmock.NewResult(1, 1))

	// dbmock.ExpectExec("UPDATE `authentication` SET").WillReturnResult(sqlmock.NewResult(1, 1))

	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// mocking token payload with userDetails as employee
	// well generally there would be all the details in this struct
	// but for conducting tests I have just populated the needful feilds
	var userdetails utility.UserDetails
	userdetails.ID = 2
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 1
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
	request := httptest.NewRequest(http.MethodPut, "/kyc", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostKycDetails(w, request)
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
	if data.Message != "Document upload successfully" || w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}
