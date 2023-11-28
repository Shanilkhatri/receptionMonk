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

// testcase #1 PutWallet with correct Struct and also correct Data,show 200.
func TestPutWalletWithCorrectStruct(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        1,
		"charge":    "opo",
		"reason":    "operuyuer",
		"cost":      54,
		"epoch":     1698214419,
		"companyId": 2,
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
	dbmock.ExpectExec("INSERT INTO `wallet`").WillReturnResult(sqlmock.NewResult(1, 1))

	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/wallets", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutWallet(w, request)
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

	if data.Message != "Wallet information added successfully." || w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #2 PutWallet with required fields,show 400.
func TestPutWalletWithRequiredFields(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        0,
		"charge":    "",
		"reason":    "",
		"cost":      54,
		"epoch":     1698214419,
		"companyId": 0,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/wallets", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutWallet(w, request)

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

	if data.Message != "Please fill the required fields properly, leaving them vacant will result in non-submission." || w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #3 PutWallet sql errors,show 400.
func TestPutWalletWithSqlErrors(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        1,
		"charge":    "opo",
		"reason":    "operuyuer",
		"cost":      54,
		"epoch":     1698214419,
		"companyId": 2,
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

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("INSERT INTO `wallet`").WillReturnError(errors.New("SQL error"))
	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/wallets", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutWallet(w, request)

	if w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #1 PostWallet with correct Struct and also correct Data,show 200.
func TestPostWalletWithCorrectStruct(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        1,
		"charge":    "opo",
		"reason":    "operuyuer",
		"cost":      54,
		"epoch":     1698214419,
		"companyId": 2,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	expectedWallet := models.Wallet{
		Id:        1,
		Charge:    "",
		Reason:    "",
		Cost:      51,
		Epoch:     1698214419,
		CompanyId: 2,
	}

	// make expected Ticket a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "charge", "reason", "cost", "epoch", "companyId"}).AddRow(expectedWallet.Id, expectedWallet.Charge, expectedWallet.Reason, expectedWallet.Cost, expectedWallet.Epoch, expectedWallet.CompanyId)

	dbmock.ExpectQuery("SELECT \\* FROM `wallet` WHERE id = ?").WillReturnRows(rows)
	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("UPDATE `wallet` SET ").WillReturnResult(sqlmock.NewResult(1, 1))

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/wallets", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostWallet(w, request)
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

	if data.Message != "Wallet information updated successfully." || w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #2 PostWallet if company is different ,show 403.
func TestPostWalletWithDiffernetCompanyId(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        1,
		"charge":    "opo",
		"reason":    "operuyuer",
		"cost":      54,
		"epoch":     1698214419,
		"companyId": 2,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 21

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	expectedWallet := models.Wallet{
		Id:        1,
		Charge:    "",
		Reason:    "",
		Cost:      51,
		Epoch:     1698214419,
		CompanyId: 2,
	}

	// make expected Ticket a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "charge", "reason", "cost", "epoch", "companyId"}).AddRow(expectedWallet.Id, expectedWallet.Charge, expectedWallet.Reason, expectedWallet.Cost, expectedWallet.Epoch, expectedWallet.CompanyId)

	dbmock.ExpectQuery("SELECT \\* FROM `wallet` WHERE id = ?").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/wallets", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostWallet(w, request)
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

	if data.Message != "You are not authorized to make this request." || w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #3 PostWallet if Empty User Struct(user struct function return false.) ,show 403.
func TestPostWalletWithEmptyUserStruct(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        1,
		"charge":    "opo",
		"reason":    "operuyuer",
		"cost":      54,
		"epoch":     1698214419,
		"companyId": 2,
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool: false,
	}

	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/wallets", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostWallet(w, request)

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

	if data.Message != "You are not authorized to make this request." || w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #4 PostWallet if User Type Is User ,show 403.
func TestPostWalletWithUserTypeIsUser(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        1,
		"charge":    "tdtrs",
		"reason":    "ttrwtertx",
		"cost":      54,
		"epoch":     1698214419,
		"companyId": 2,
	}

	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required

	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user"
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}

	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/wallets", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostWallet(w, request)

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

	if data.Message != "You are not authorized to make this request." || w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #1 GetWallet with correct Struct and also correct Data,show 200.
func TestGetWalletWithCorrectStruct(t *testing.T) {

	queryParams := map[string]string{
		"id":        "1",
		"companyId": "2",
	}

	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	expectedWallet := models.Wallet{
		Id:        1,
		Charge:    "",
		Reason:    "",
		Cost:      51,
		Epoch:     1698214419,
		CompanyId: 2,
	}

	// make expected Ticket a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "charge", "reason", "cost", "epoch", "companyId"}).AddRow(expectedWallet.Id, expectedWallet.Charge, expectedWallet.Reason, expectedWallet.Cost, expectedWallet.Epoch, expectedWallet.CompanyId)

	dbmock.ExpectQuery("SELECT \\* FROM `wallet` ").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB
	url := "/wallets?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetWallet(w, request)
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

	if data.Message != "Fetched wallet transactions successfully." || w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #2 GetWallet with user struct function return false,show 403.
func TestGetWalletWithIncorrectUserStruct(t *testing.T) {

	queryParams := map[string]string{
		"id":        "1",
		"companyId": "2",
	}

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool: false,
	}

	url := "/wallets?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetWallet(w, request)

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

	if data.Message != "You are not authorized to make this request." || w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #3 GetWallet with AccountType user,show 403.
func TestGetWalletWithAccountTypeUser(t *testing.T) {

	queryParams := map[string]string{
		"id":        "1",
		"companyId": "2",
	}

	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user" //set account type is user.
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}

	url := "/wallets?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetWallet(w, request)

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

	if data.Message != "Unauthorized access! You are not authorized to make this request." || w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #4 GetWallet different company id,show 403.
func TestGetWalletWithDifferentCompanyId(t *testing.T) {

	queryParams := map[string]string{
		"id":        "1",
		"companyId": "2",
	}

	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 45 //set company id.

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}

	url := "/wallets?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetWallet(w, request)

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

	if data.Message != "Unauthorized access! You are not authorized to make this request." || w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #5 GetWallet with Sql return error "no rows found" error,show 500.
func TestGetWalletWithSqlErrors(t *testing.T) {

	queryParams := map[string]string{
		"id":        "1",
		"companyId": "2",
	}

	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM `wallet` ").WillReturnError(sql.ErrNoRows) // Simulate a "no rows found" error.
	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB
	url := "/wallets?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetWallet(w, request)
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

	if data.Message != "Internal server error, Any serious issues which cannot be recovered from." || w.Result().StatusCode != 500 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// testcase #6 GetWallet withNo result were found ,show 200.
func TestGetWalletWithNOResultFounds(t *testing.T) {

	queryParams := map[string]string{
		"id":        "1",
		"companyId": "2",
	}

	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// make expected Ticket a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "charge", "reason", "cost", "epoch", "companyId"})

	dbmock.ExpectQuery("SELECT \\* FROM `wallet` ").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB
	url := "/wallets?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetWallet(w, request)
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

	if data.Message != "No result were found for this search." || w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// Test#1 wallet delete success ,if wallet id not empty or nil,show 200.
func TestWalletDeleteSuccess(t *testing.T) {
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	expectedWallet := models.Wallet{
		Id:        1,
		Charge:    "",
		Reason:    "",
		Cost:      51,
		Epoch:     1698214419,
		CompanyId: 2,
	}

	// make expected wallet a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "charge", "reason", "cost", "epoch", "companyId"}).AddRow(expectedWallet.Id, expectedWallet.Charge, expectedWallet.Reason, expectedWallet.Cost, expectedWallet.Epoch, expectedWallet.CompanyId)

	dbmock.ExpectQuery("SELECT \\* FROM `wallet` ").WillReturnRows(rows)

	// I  expect a Delete Query execution and for that :
	dbmock.ExpectExec("DELETE FROM `wallet` WHERE id = ?").WillReturnResult(sqlmock.NewResult(1, 1))

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set walletId to be deleted to 1
	walletIdToDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/wallet?" + "id=" + walletIdToDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteWallet(w, request)

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

	if data.Message != "Wallet deleted successfully." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// Test#2 check wallet id validation for delete,show 400.
func TestWalletDeleteValidation(t *testing.T) {

	walletIdToDel := "" //if id is empty string.
	// here we will prepare the url with parameters to pass to our request
	url := "/wallet?" + "id=" + walletIdToDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteWallet(w, request)

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

// Test#3 wallet delete if user struct function return false,show 403.
func TestWalletDeleteUserStruct(t *testing.T) {

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool: false,
	}

	// set walletId to be deleted to 1
	walletIdToDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/wallet?" + "id=" + walletIdToDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteWallet(w, request)

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
	log.Println(data.Message)
	if data.Message != "You are not authorized to make this request." && w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// Test#4 wallet delete if User Account type is user,show 403.
func TestWalletDeleteUserAccountTypeIsUser(t *testing.T) {
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user"
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}

	// set walletId to be deleted to 1
	walletIdToDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/wallet?" + "id=" + walletIdToDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteWallet(w, request)

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

// Test#5 wallet return empty rows ,show 400.
func TestWalletDeleteReturnEmptyRows(t *testing.T) {
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// make expected wallet a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "charge", "reason", "cost", "epoch", "companyId"})

	dbmock.ExpectQuery("SELECT \\* FROM `wallet` ").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set walletId to be deleted to 1
	walletIdToDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/wallet?" + "id=" + walletIdToDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteWallet(w, request)

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

	if data.Message != "no rows in result set" && w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// Test#6 wallet Delete return errors no rows in result set,show 400.
func TestWalletDeleteReturnSqlErrors(t *testing.T) {
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	expectedWallet := models.Wallet{
		Id:        1,
		Charge:    "",
		Reason:    "",
		Cost:      51,
		Epoch:     1698214419,
		CompanyId: 2,
	}

	// make expected wallet a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "charge", "reason", "cost", "epoch", "companyId"}).AddRow(expectedWallet.Id, expectedWallet.Charge, expectedWallet.Reason, expectedWallet.Cost, expectedWallet.Epoch, expectedWallet.CompanyId)

	dbmock.ExpectQuery("SELECT \\* FROM `wallet` ").WillReturnRows(rows)

	// I  expect a Delete Query execution and for that :
	dbmock.ExpectExec("DELETE FROM `wallet` WHERE id = ?").WillReturnError(sql.ErrNoRows)
	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set walletId to be deleted to 1
	walletIdToDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/wallet?" + "id=" + walletIdToDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteWallet(w, request)

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

	if data.Message != "no rows in result set" && w.Result().StatusCode != 400 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// Test#7 wallet delete ,if companyId is different,show 403.
func TestWalletDeleteDifferentCompanyId(t *testing.T) {
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 2

	// Mocking the utility functions that are used there
	Helper = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	expectedWallet := models.Wallet{
		Id:        1,
		Charge:    "",
		Reason:    "",
		Cost:      51,
		Epoch:     1698214419,
		CompanyId: 25,
	}

	// make expected wallet a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "charge", "reason", "cost", "epoch", "companyId"}).AddRow(expectedWallet.Id, expectedWallet.Charge, expectedWallet.Reason, expectedWallet.Cost, expectedWallet.Epoch, expectedWallet.CompanyId)

	dbmock.ExpectQuery("SELECT \\* FROM `wallet` ").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set walletId to be deleted to 1
	walletIdToDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/wallet?" + "id=" + walletIdToDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	DeleteWallet(w, request)

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

	if data.Message != "You are not authorized to make this request." && w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}
