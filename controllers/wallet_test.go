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
	Utility = MockHelper{
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
	request := httptest.NewRequest(http.MethodPut, "/wallets", bytes.NewBuffer(requestBody))
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
	Utility = MockHelper{
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
	request := httptest.NewRequest(http.MethodPut, "/wallets", bytes.NewBuffer(requestBody))
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
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool: false,
	}

	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/wallets", bytes.NewBuffer(requestBody))
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
	Utility = MockHelper{
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}

	// Create a mock Request
	request := httptest.NewRequest(http.MethodPut, "/wallets", bytes.NewBuffer(requestBody))
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
