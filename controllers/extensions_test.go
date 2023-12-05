package controllers

import (
	// Standard Library Packages
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reakgo/models"
	"reakgo/utility"
	"testing"

	// Third Party Packages
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Tests the add extension function by sending incorrect parameters to see if we get proper response
func TestExtensionAddWithIncorrectData(t *testing.T) {
	// Create a mock HTTP request payload
	jsonData := `{
	"id": 1,
  	"first_name": "Jeanette",
  	"last_name": "Penddreth",
  	"email": "jpenddreth0@census.gov",
  	"gender": "Female",
  	"ip_address": "26.58.193.2"
}`

	// Writing the payload and request attributes
	r := httptest.NewRequest("POST", "/extension/add", bytes.NewBuffer([]byte(jsonData)))

	// Create a mock HTTP response writer
	w := httptest.NewRecorder()

	// Call the target Extension, here we don't need to mock anything as it should fail before all this
	_ = PostExtension(w, r)

	// Assertions, Check if 400 is returned as payload is incorrect
	if w.Result().Status != fmt.Sprint(400) {
		t.Errorf("Expected controller to return 400, got %s", w.Result().Status)
	}
}

// Tests the add extension function by sending correct parameters to see if we get proper response
func TestExtensionAddWithCorrectData(t *testing.T) {
	// Create a mock HTTP request payload
	jsonData := `{
		"id": 2,
		"extension": "sfg23445",
		"userid": 1,
		"department": 3,
		"sipserver": "sip.example.com",
		"sipusername": "user1123232",
		"sippassword": "123",
		"sipport": "506090"
	}`

	// DB Wizardry, We need to hook utility.Db to have sqlx cursor, So we mock the entire DB set
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// All of the steps need to be provided to mock, Since we're using MustBegin hence the ExpectBegin parameter to prepare the mock
	dbmock.ExpectBegin()

	// Expect Update to Extensions (Not sure why we're updating it in add extension but fuck it)
	dbmock.ExpectExec("UPDATE `extensions`").WillReturnResult(sqlmock.NewResult(1, 1))

	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	Helper = MockHelper{
		MockReturnUserDetailsResult: nil,
	}

	r := httptest.NewRequest("POST", "/extension/add", bytes.NewBuffer([]byte(jsonData)))

	// Create a mock HTTP response writer
	w := httptest.NewRecorder()

	_ = PostExtension(w, r)

	if w.Result().Status != "200 OK" {
		t.Errorf("Expected controller to return 200, got %s", w.Result().Status)
	}
	t.Errorf("Expected controller to return 200, got %s", w.Result().Status)
}

// TEST #1 - PostExtension with correct Struct and also correct Data,show 200.
func TestExtensionPostSuccess(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":          1,
		"extension":   "extesionitem",
		"userid":      1,
		"department":  1,
		"sipserver":   "newserver",
		"sipusername": "sipUser",
		"sippassword": "123",
		"sipport":     "2456",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 1

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()

	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()

	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("UPDATE `extensions` INNER JOIN `authentication` ON `extensions`.`user_id` = `authentication`.`id` INNER JOIN `company` ON `authentication`.`company_id` = `company`.`id` SET").WillReturnResult(sqlmock.NewResult(1, 1))

	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockReturnUserDetailsResult:               nil,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/extension", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostExtension(w, request)
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

	if data.Message != "Extension updated successfully." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 - PostExtension with incorrect user Struct(token payload function return false) ,show 400.
func TestExtensionPostReturnError(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":          1,
		"extension":   "extesionitem",
		"userid":      1,
		"department":  1,
		"sipserver":   "newserver",
		"sipusername": "sipUser",
		"sippassword": "123",
		"sipport":     "2456",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                   nil,
		MockCheckTokenPayloadAndReturnUserBool: false,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/extension", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostExtension(w, request)

	if w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #3 - PostExtension with Required Fields show 400.
func TestExtensionPostRequiredFields(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":          0, // incorrect extension id .
		"extension":   "extesionitem",
		"userid":      0, // incorrect userId.
		"department":  1,
		"sipserver":   "newserver",
		"sipusername": "sipUser",
		"sippassword": "123",
		"sipport":     "2456",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 1

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockReturnUserDetailsResult:               nil,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/extension", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostExtension(w, request)

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

	if data.Message != "Bad request, Incorrect payload or call." && w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #1 - PutExtension with correct Struct and also correct Data and show 200.
func TestExtensionPutSuccess(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":          1,
		"extension":   "extesionitem",
		"userid":      1,
		"department":  1,
		"sipserver":   "newserver",
		"sipusername": "sipUser",
		"sippassword": "123",
		"sipport":     "2456",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user"
	userdetails.CompanyID = 2

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()

	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()

	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("INSERT INTO `extension`").WillReturnResult(sqlmock.NewResult(1, 1))

	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockReturnUserDetailsResult:               nil,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/extension", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutExtension(w, request)
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

	if data.Message != "Extension added successfully." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 - PutExtension with incorrect user Struct(user return error or false), show 400.
func TestExtensionPutReturnError(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":          1,
		"extension":   "extesionitem",
		"userid":      1,
		"department":  1,
		"sipserver":   "newserver",
		"sipusername": "sipUser",
		"sippassword": "123",
		"sipport":     "2456",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                   nil,
		MockCheckTokenPayloadAndReturnUserBool: false,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/extension", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutExtension(w, request)

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

	if data.Message != "Bad request, Incorrect payload or call." && w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #3 - PutExtension with Required Fields,show 400.
func TestExtensionPutRequiredFields(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":          1,
		"extension":   "extesionitem",
		"userid":      1,
		"department":  0, //inccorect department
		"sipserver":   "newserver",
		"sipusername": "sipUser",
		"sippassword": "123",
		"sipport":     "2456",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user"
	userdetails.CompanyID = 1

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockReturnUserDetailsResult:               nil,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/extension", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutExtension(w, request)

	if w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #4 - PutExtension with check sql error.
func TestExtensionPutSqlErrorCheck(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":          1,
		"extension":   "extesionitem",
		"userid":      1,
		"department":  1,
		"sipserver":   "newserver",
		"sipusername": "sipUser",
		"sippassword": "123",
		"sipport":     "2456",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user"
	userdetails.CompanyID = 1

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()

	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	dbmock.ExpectBegin()

	expectedErr := mysql.MySQLError{Number: 1062, Message: "duplicate entry"}
	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("INSERT INTO `extension`").
		WillReturnError(&expectedErr)
	dbmock.ExpectCommit()
	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockReturnUserDetailsResult:               nil,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/extension", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutExtension(w, request)

	if w.Result().StatusCode != 500 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// Test#1 Extension delete success ,if extension id not empty or nil,show 200.
func TestExtensionDeleteSuccess(t *testing.T) {

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println("mockErr: ", err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I  expect a Delete Query execution and for that :
	dbmock.ExpectExec("DELETE FROM extension WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set extensionId to be deleted to 1
	extensionIdtoDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/extension?" + "id=" + extensionIdtoDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteExtension(w, request)

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

	if data.Message != "Extension deleted successfully." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// Test#2 check extension id validation for delete,show 400.
func TestExtensionDeleteValidation(t *testing.T) {

	// set extension id
	extensionIdtoDel := ""
	// here we will prepare the url with parameters to pass to our request
	url := "/extension?" + "id=" + extensionIdtoDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteExtension(w, request)

	if w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// Test#3 check sql error foreign key constraint violation(1451)..
func TestExtensionSqlErrorCheck(t *testing.T) {

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println("mockErr: ", err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I  expect a Delete Query execution and for that :
	// dbmock.ExpectExec("DELETE FROM extension WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))
	expectedErr := mysql.MySQLError{Number: 1451, Message: "Foreign key constraint violation."}
	dbmock.ExpectExec("DELETE FROM extension WHERE id = ?").
		WillReturnError(&expectedErr)
	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set extensionId to be deleted to 1
	extensionIdtoDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/extension?" + "id=" + extensionIdtoDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteExtension(w, request)

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

	if data.Message != "Foreign key constraint violation." && w.Result().StatusCode != 500 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #1 - GetExtension with correct Struct and also correct Data,succes 200.
func TestExtensionGetSuccess(t *testing.T) {

	queryParams := map[string]string{
		"id":        "1",
		"companyId": "1",
	}

	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user"
	userdetails.CompanyID = 1

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()

	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	expectedExtension := models.Extensions{
		Id:          1,
		Extension:   "extension",
		UserId:      1,
		Department:  1,
		SipServer:   "servername",
		SipUserName: "sipUser@name",
		SipPassword: "123",
		SipPort:     "2089",
	}

	// make expected user a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "extension", "userId", "department", "sipserver", "sipusername", "sippassword", "sipport"}).AddRow(expectedExtension.Id, expectedExtension.Extension, expectedExtension.UserId, expectedExtension.Department, expectedExtension.SipServer, expectedExtension.SipUserName, expectedExtension.SipPassword, expectedExtension.SipPort)

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM extension").WillReturnRows(rows)
	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockReturnUserDetailsResult:               nil,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// here we will prepare the url with parameters to pass to our request
	url := "/extension?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetExtensionData(w, request)

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

	if data.Message != "Returns all matching  Extensions." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 if user (token payload function return false), show 400.
func TestGetExtensionUserIdEmpty(t *testing.T) {
	queryParams := map[string]string{
		"id":        "1",
		"companyId": "",
	}

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                   nil,
		MockCheckTokenPayloadAndReturnUserBool: false,
	}

	// here we will prepare the url with parameters to pass to our request
	url := "/extension?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetExtensionData(w, request)

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

	if data.Message != "Bad request, Incorrect payload or call." && w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #3 - GetExtension return no result found,show 500.
func TestExtensionReturnSqlErrors(t *testing.T) {

	queryParams := map[string]string{
		"id":        "1",
		"companyId": "1",
	}

	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user"
	userdetails.CompanyID = 1

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()

	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM extension Where 1=1 AND id=\\? AND type IN \\('user'\\)").WillReturnError(sql.ErrNoRows) // Simulate a "no rows found" error.

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockReturnUserDetailsResult:               nil,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// here we will prepare the url with parameters to pass to our request
	url := "/extension?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetExtensionData(w, request)

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

// TEST #4 if company id is different, show 403.
func TestGetExtensionCompanyIdEmpty(t *testing.T) {
	queryParams := map[string]string{
		"id":        "1",
		"companyId": "2", //incorrect company id
	}

	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user"
	userdetails.CompanyID = 4

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		MockSessionGetResult:                      nil,
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockReturnUserDetailsResult:               nil,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}

	// open Mock DB connection
	mockDB, _, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/extension?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetExtensionData(w, request)

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

	if data.Message != "You are not authorized for this request." && w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}
