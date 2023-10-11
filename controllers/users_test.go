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
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MockResponseWriter struct {
	Status  int
	Headers http.Header
	Body    []byte
}

func (m *MockResponseWriter) Header() http.Header {
	if m.Headers == nil {
		m.Headers = make(http.Header)
	}
	return m.Headers
}

func (m *MockResponseWriter) Write(b []byte) (int, error) {
	m.Body = append(m.Body, b...)
	return len(b), nil
}

func (m *MockResponseWriter) WriteHeader(statusCode int) {
	m.Status = statusCode
}

// TEST #1 - putUser with correct Struct and also correct Data
func TestUserPutWithCorrectData(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":                    0,
		"name":                  "hguhduhs",
		"email":                 "user@example.com",
		"passwordHash":          "1234",
		"twoFactorKey":          "55",
		"twoFactorRecoveryCode": "59898",
		"dob":                   "2023-10-05",
		"accountType":           "owner",
		"companyId":             0,
		"status":                "active",
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
	dbmock.ExpectExec("INSERT INTO `authentication`").WillReturnResult(sqlmock.NewResult(1, 1))

	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                   nil,
		MockCheckTokenPayloadAndReturnUserBool: false,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutUser(w, request)
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

	if data.Message != "User created successfully." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 - putUser with incorrect Date string
// correct date string is in format "yyyy-mm-dd"

// I have also added a check for incorrect email
// try replacing email with - userexample.com
func TestUserPutWithIncorrectDateString(t *testing.T) {
	jsonData := map[string]interface{}{
		"id":                    0,
		"name":                  "hguhduhs",
		"email":                 "user@example.com",
		"passwordHash":          "1234",
		"twoFactorKey":          "55",
		"twoFactorRecoveryCode": "59898",
		"dob":                   "2023-20-05", //incorrect date
		"accountType":           "owner",
		"companyId":             0,
		"status":                "active",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// you can notice that as we are expecting an error even before touching the DB ops,
	// so we haven't mocked DB in this case

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                   nil,
		MockCheckTokenPayloadAndReturnUserBool: false,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutUser(w, request)

	// we can even read body using io package and test even specific messages
	// for now I have skipped it, might be added in future

	if w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Result().StatusCode)
	}
}

// TEST #3 - putUser with faulty structure/incomplete data
// you can try by removing any key-val from json data

func TestUserPutWithFaultyStruct(t *testing.T) {
	jsonData := map[string]interface{}{
		"id":                    0,
		"name":                  "hguhduhs",
		"email":                 "user@example.com",
		"passwordHash":          "1234",
		"twoFactorKey":          "55",
		"twoFactorRecoveryCode": "59898",
		"dob":                   "2023-10-05", //incorrect date
		"accountType":           "owner",
		"companyId":             0,
		// there is no "status"
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// you can notice that as we are expecting an error even before touching the DB ops,
	// so we haven't mocked DB in this case

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                   nil,
		MockCheckTokenPayloadAndReturnUserBool: false,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutUser(w, request)

	// we can even read body using io package and test even specific messages
	// for now I have skipped it, might be added in future

	if w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Result().StatusCode)
	}
}

// TEST #4 - putUser - when a guest tries to register as user
// yup! guests cannot signup as users, only owners can.

func TestUserPutGuestRegAsUser(t *testing.T) {
	jsonData := map[string]interface{}{
		"id":                    0,
		"name":                  "hguhduhs",
		"email":                 "user@example.com",
		"passwordHash":          "1234",
		"twoFactorKey":          "55",
		"twoFactorRecoveryCode": "59898",
		"dob":                   "2023-10-05",
		"accountType":           "user", // nothing wrong, just not allowed!
		"companyId":             0,
		"status":                "active",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// you can notice that as we are expecting an error even before touching the DB ops,
	// so we haven't mocked DB in this case

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                   nil,
		MockCheckTokenPayloadAndReturnUserBool: false,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutUser(w, request)

	// we can even read body using io package and test even specific messages
	// for now I have skipped it, might be added in future

	if w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Result().StatusCode)
	}
}

// TEST #5 - putUser - if user registers with owners token he can register
// we call it - registration through invite links(owners token set in header)
// we have bypassed the checkACL thats why we have to set tokenPayload in header
// -> we have to mock sessionget to give either user,owner or super-admin
// -> we also have to mock tokenPayload to give us userdetails

func TestUserPutWithOwnersToken(t *testing.T) {
	jsonData := map[string]interface{}{
		"id":                    0,
		"name":                  "hguhduhs",
		"email":                 "user@example.com",
		"passwordHash":          "1234",
		"twoFactorKey":          "55",
		"twoFactorRecoveryCode": "59898",
		"dob":                   "2023-10-05",
		"accountType":           "user",
		"companyId":             0,
		"status":                "active",
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

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("INSERT INTO `authentication`").WillReturnResult(sqlmock.NewResult(1, 1))

	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      "owner", //setting session also as owner
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutUser(w, request)
	err = dbmock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expectations were not met %s", err)
	}

	// we can even read body using io package and test even specific messages
	// for now I have skipped it, might be added in future

	if w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// -------------------USER POST TESTS---------------------

// -> firstly, when anyone posting update they'll go through an ACL check which gives us
// thier details that we have in the DB through cache or through DB itself
// -> here we don't have checkACL as we totally bypass it so the details part will be mocked by us.
// -> There are three checks that occur when under controller:
//		- tokenPayload which if not present that means user is a guest and he won't get access
//		- if userStruct.id == 0 (case if we misplaced id while transfer so the POST won't get through)
//		- lastly if userStruct.id != userDetails.id (the POST won't get through as each user can update only their details)
// based on the above info and to cross check if these conditions really work, let's begin testing!

// TEST #1 postUser with correct data and struct
