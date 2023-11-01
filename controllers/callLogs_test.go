package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reakgo/models"
	"reakgo/utility"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// TEST #1 - putCallLogs with correct Struct and also correct Data
// -> To put callLogs succesfully the user should be an owner
// -> we'll mock the user details from the token and make accountType : owner
// -> we'll also mock the Db to perform ops of insertion
// expecting a 200
func TestCallLogsPutWithCorrectData(t *testing.T) {
	// callLogs data to be Put
	jsonData := map[string]interface{}{
		"id":            0,
		"callFrom":      "abc",
		"callTo":        "def",
		"callPlacedAt":  "ghi",
		"callDuration":  "jkl",
		"callExtension": "mno",
		"companyId":     1,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as owner
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 3
	userdetails.AccountType = "owner" // type set to owner

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("INSERT INTO callLogs").WillReturnResult(sqlmock.NewResult(1, 1))

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
	request := httptest.NewRequest(http.MethodPut, "/calllogs", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutCallLogs(w, request)
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

	if data.Message != "CallLog created successfully." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 - putCallLogs with incorrect Data type
// -> To put callLogs succesfully the user should be an owner
// -> we'll mock the user details from the token and make accountType : owner
// -> we'll not mock Db as we expect to catch an error before DB ops
// expecting a 400
func TestCallLogsPutWithIncorrectData(t *testing.T) {
	// callLogs data to be Put
	jsonData := map[string]interface{}{
		"id":            0,
		"callFrom":      15, // incorrect value type
		"callTo":        "def",
		"callPlacedAt":  "ghi",
		"callDuration":  "jkl",
		"callExtension": "mno",
		"companyId":     1,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as owner
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 3
	userdetails.AccountType = "owner" // type set to owner

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                   nil,
		MockCheckTokenPayloadAndReturnUserBool: false,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/calllogs", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutCallLogs(w, request)

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
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #3 - putCallLogs with unauthorized user type
// -> To put callLogs succesfully the user should be an owner
// -> we'll mock the user details from the token and make accountType : user
// -> we'll not mock Db as we expect to catch an error before DB ops
// expecting a 403
func TestCallLogsPutWithUnAuthUserType(t *testing.T) {
	// callLogs data to be Put
	jsonData := map[string]interface{}{
		"id":            0,
		"callFrom":      "abc",
		"callTo":        "def",
		"callPlacedAt":  "ghi",
		"callDuration":  "jkl",
		"callExtension": "mno",
		"companyId":     1,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as owner
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 3
	userdetails.AccountType = "user" // unauthorized type user

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                   nil,
		MockCheckTokenPayloadAndReturnUserBool: false,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/calllogs", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutCallLogs(w, request)

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

	if data.Message != "You are not authorized to make this request" || w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// test #1 getCallLogs with correct Struct and also correct Data,show 200.
func TestCallLogsGettWithCorrectData(t *testing.T) {
	// callLogs data to be Put
	queryParams := map[string]string{
		"id":        "1",
		"companyId": "0",
	}

	// mocking token payload with userDetails as owner
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner" // type set to owner

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// mocking the user detail I expect from the Db op
	expectedCallLogs := models.CallLogs{
		Id:            1,
		CallFrom:      "A",
		CallTo:        "B",
		CallPlacedAt:  "jbl",
		CallDuration:  "12",
		CallExtension: "rtw",
		CompanyId:     2,
	}
	// make expected user a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "callform", "callto", "callplacedat", "callduration", "callextension", "companyId"}).AddRow(expectedCallLogs.Id, expectedCallLogs.CallFrom, expectedCallLogs.CallTo, expectedCallLogs.CallPlacedAt, expectedCallLogs.CallDuration, expectedCallLogs.CallExtension, expectedCallLogs.CompanyId)

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM `calllogs` WHERE 1=1 AND id=?").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/calllogs?" + buildQuery(queryParams)

	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetCallLogsDetails(w, request)

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

	if data.Message != "Returns all matching  calllogs." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// test #2 getCallLogs if empty user struct,show 403.
func TestCallLogsGettWithEmptyuserStruct(t *testing.T) {
	// callLogs data to be Put
	queryParams := map[string]string{
		"id":        "1",
		"companyId": "0",
	}

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool: false,
	}

	// here we will prepare the url with parameters to pass to our request
	url := "/calllogs?" + buildQuery(queryParams)

	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetCallLogsDetails(w, request)

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

	if data.Message != "You are not authorized to make this request." && w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// test #3 getCallLogs different companyId ,show 403.
func TestCallLogsGetWithDifferentCompanyId(t *testing.T) {
	// callLogs data to be Put
	queryParams := map[string]string{
		"id":        "0",
		"companyId": "7", //"user has a companyId of 3, but I search for a different companyId, which is 7.".
	}

	// mocking token payload with userDetails as owner
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner" // type set to owner
	userdetails.CompanyID = 3

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}

	// here we will prepare the url with parameters to pass to our request
	url := "/calllogs?" + buildQuery(queryParams)

	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetCallLogsDetails(w, request)

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

	if data.Message != "Unauthorized access! You are not authorized to make this request." && w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// test #4 getCallLogs with user types is user,show 403.
func TestCallLogsGettWithUserTypeIsUser(t *testing.T) {
	// callLogs data to be Put
	queryParams := map[string]string{
		"id":        "1",
		"companyId": "0",
	}

	// mocking token payload with userDetails as owner
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user" // type set to owner

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}

	// here we will prepare the url with parameters to pass to our request
	url := "/calllogs?" + buildQuery(queryParams)

	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetCallLogsDetails(w, request)

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

	if data.Message != "Unauthorized access! You are not authorized to make this request." && w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// test #5 getCallLogs sql return errors (No result were found for this search.) ,show 200.
func TestCallLogsGettWithSqlErrors(t *testing.T) {
	// callLogs data to be Put
	queryParams := map[string]string{
		"id":        "1",
		"companyId": "0",
	}

	// mocking token payload with userDetails as owner
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner" // type set to owner

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
	rows := sqlmock.NewRows([]string{"id", "callform", "callto", "callplacedat", "callduration", "callextension", "companyId"})
	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM `calllogs` WHERE 1=1 AND id=?").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/calllogs?" + buildQuery(queryParams)

	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetCallLogsDetails(w, request)

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

	if data.Message != "No result were found for this search." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// test #6 getCallLogs sql return errors ,show 500.
func TestGetCallLogsSqlError(t *testing.T) {
	// callLogs data to be Put
	queryParams := map[string]string{
		"id":        "1",
		"companyId": "0",
	}

	// mocking token payload with userDetails as owner
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner" // type set to owner

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM `calllogs` WHERE 1=1 AND id=?").WillReturnError(sql.ErrNoRows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/calllogs?" + buildQuery(queryParams)

	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetCallLogsDetails(w, request)

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
