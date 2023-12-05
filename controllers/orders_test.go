package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reakgo/models"
	"reakgo/utility"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Function to build a query string from a map of query parameters to make it easy for us
func buildQuery(params map[string]string) string {
	var query string
	for key, value := range params {
		query += key + "=" + value + "&"
	}
	return strings.TrimRight(query, "&")
}

// TEST #1 - GetOrders with correct parameters
// -> There are three access levels based on "accountType" to get orders info and they are:
// -> when AccountType : user -> this account can only get orders related to his userId
// -> when AccountType : owner -> this account can get orders related to his userId as well as all users working in his company
// -> when AccountType : super-admin -> this account can get orders info related to anyone(highest level of access rights)
// -> for the first test we'll mock the user details from the token and make accountType : user
// -> we'll also mock the Db to perform ops of fetch
// expecting a 200
func TestGetOrdersWithTypeUser(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
	// instead of sending json data we have to send parametrs in query
	queryParams := map[string]string{
		"id":        "",
		"date_from": "",
		"date_to":   "",
		"companyId": "",
		"userId":    "1", // getting data for userId #1
	}
	// mocking token payload with userDetails as user
	// haven't populated all the fields as they aren't required
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user" // type set to user
	userdetails.CompanyID = 1

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

	// mocking the user detail I expect from the Db op
	expectedOrders := models.Orders{
		Id:        1,
		ProductId: 1,
		PlacedOn:  1696932888,
		Expiry:    1698747288,
		Price:     250,
		Buyer:     1, // This is basically userId
		Status:    "unpaid",
	}
	// make expected user a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "productId", "placedOn", "expiry", "price", "buyer", "status"}).AddRow(expectedOrders.Id, expectedOrders.ProductId, expectedOrders.PlacedOn, expectedOrders.Expiry, expectedOrders.Price, expectedOrders.Buyer, expectedOrders.Status)

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT orders.id,orders.productId").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/orders?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetOrders(w, request)
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

// TEST #2 - GetOrders with correct parameters
// -> user tries to access someone elses orders
// -> for the second test we'll mock the user details from the token and make accountType : user
// -> we won't mock DB as we expect to get error unauhtorized access before the DB ops
// expecting a 403
func TestGetOrdersWithTypeUserAccessingAnotherUser(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
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
		// MockStrToInt64Int:                         2,
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
	GetOrders(w, request)

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

// TEST #3 - GetOrders with correct parameters
// -> owner tries to access someone else's orders of diff company
// -> for the third test we'll mock the user details from the token and make accountType : owner
// -> we will mock DB as we expect to get result(length)=0
// expecting a 400
func TestGetOrdersWithTypeOwnerAccessingAnotherUserOfDiffCompany(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
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

	// make orders row that will be returned but empty
	rows := sqlmock.NewRows([]string{"id", "productId", "placedOn", "expiry", "price", "buyer", "status"})

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT orders.id,orders.productId").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/orders?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetOrders(w, request)

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

// TEST #4 - GetOrders with correct parameters
// -> owner tries to access someone else's orders of same company
// -> for the fourth test we'll mock the user details from the token and make accountType : owner
// -> we will mock DB as we expect to get result
// expecting a 200
func TestGetOrdersWithTypeOwnerAccessingAnotherUserOfSameCompany(t *testing.T) {
	// parameters to be send on the basis of which we can filter orders and fetch
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

	// mocking the user detail I expect from the Db op
	expectedOrders := models.Orders{
		Id:        1,
		ProductId: 1,
		PlacedOn:  1696932888,
		Expiry:    1698747288,
		Price:     250,
		Buyer:     2, // This is basically userId
		Status:    "unpaid",
	}
	// make expected user a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "productId", "placedOn", "expiry", "price", "buyer", "status"}).AddRow(expectedOrders.Id, expectedOrders.ProductId, expectedOrders.PlacedOn, expectedOrders.Expiry, expectedOrders.Price, expectedOrders.Buyer, expectedOrders.Status)

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT orders.id,orders.productId").WillReturnRows(rows)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// here we will prepare the url with parameters to pass to our request
	url := "/orders?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	GetOrders(w, request)

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

// TEST #1 - putOrder with correct Struct and also correct Data.
func TestOrderPutSuccess(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        1,
		"productId": 1,
		"placedOn":  1697632642,
		"expiry":    1697632642,
		"price":     55,
		"buyer":     59898,
		"status":    "paid",
		"dateFrom":  1697632642,
		"DateTo":    1698237442,
		"UserId":    1,
		"CompanyId": 21,
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

	int64Value := int64(42)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(int64Value)
	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()

	dbmock.ExpectQuery("SELECT `plan_validity` FROM `products` WHERE  productId=?").WillReturnRows(rows)

	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("INSERT INTO `orders`").WillReturnResult(sqlmock.NewResult(1, 1))

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
	request := httptest.NewRequest(http.MethodPut, "/order", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutOrder(w, request)
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

	if data.Message != "Order added successfully." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 - putOrder with incorrect user Struct(user return error or false).
func TestOrderPutWithUserReturnError(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        1,
		"productId": 1,
		"placedOn":  1697632642,
		"expiry":    1697632642,
		"price":     55,
		"buyer":     59898,
		"status":    "paid",
		"dateFrom":  1697632642,
		"DateTo":    1698237442,
		"UserId":    1,
		"CompanyId": 21,
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
	request := httptest.NewRequest(http.MethodPut, "/order", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutOrder(w, request)

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

// TEST #3 - putOrder with Required Fields,show 400.
func TestOrderPutRequiredFields(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        0, //incorrect value
		"productId": 0, //incorrect value
		"placedOn":  1697632642,
		"expiry":    1697632642,
		"price":     55,
		"buyer":     59898,
		"status":    "paid",
		"dateFrom":  1697632642,
		"DateTo":    1698237442,
		"UserId":    1,
		"CompanyId": 21,
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
	request := httptest.NewRequest(http.MethodPut, "/order", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PutOrder(w, request)

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

// TEST #1 - PostOrder with correct Struct and also correct Data.
func TestOrderPostSuccess(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        1,
		"productId": 1,
		"placedOn":  1697632642,
		"expiry":    1697632642,
		"price":     55,
		"buyer":     59898,
		"status":    "paid",
		"dateFrom":  1697632642,
		"DateTo":    1698237442,
		"UserId":    1,
		"CompanyId": 21,
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

	int64Value := int64(42)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(int64Value)
	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()

	dbmock.ExpectQuery("SELECT `plan_validity` FROM `products` WHERE  productId=?").WillReturnRows(rows)

	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("UPDATE `orders` SET").WillReturnResult(sqlmock.NewResult(1, 1))

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
	request := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostOrder(w, request)
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

	if data.Message != "Order Update successfully." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 - PostOrder with incorrect user Struct(user return error or false),show 400.
func TestOrderPostReturnError(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        1,
		"productId": 1,
		"placedOn":  1697632642,
		"expiry":    1697632642,
		"price":     55,
		"buyer":     59898,
		"status":    "paid",
		"dateFrom":  1697632642,
		"DateTo":    1698237442,
		"UserId":    1,
		"CompanyId": 21,
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
	request := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostOrder(w, request)

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

// TEST #3 - PostOrder with Required Fields ,show 400.
func TestOrderPostRequiredFields(t *testing.T) {

	jsonData := map[string]interface{}{
		"id":        0, //incorrect value
		"productId": 0, //incorrect value
		"placedOn":  1697632642,
		"expiry":    1697632642,
		"price":     55,
		"buyer":     59898,
		"status":    "paid",
		"dateFrom":  1697632642,
		"DateTo":    1698237442,
		"UserId":    1,
		"CompanyId": 21,
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
	request := httptest.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostOrder(w, request)

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

// Test1 order delete success ,if order id not empty or nil,show 200.
func TestOrderDeleteSuccess(t *testing.T) {

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println("mockErr: ", err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I  expect a Delete Query execution and for that :
	dbmock.ExpectExec("DELETE FROM orders").WillReturnResult(sqlmock.NewResult(1, 1))

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set orderId to be deleted to 1
	orderIdtoDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/order?" + "id=" + orderIdtoDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	OrderDelete(w, request)

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

	if data.Message != "Order deleted successfully." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// Test2 check order id validation for delete,show 400.
func TestOrderDeleteValidation(t *testing.T) {

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println("mockErr: ", err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I  expect a Delete Query execution and for that :
	dbmock.ExpectExec("DELETE FROM orders").WillReturnResult(sqlmock.NewResult(1, 1))

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set order id empty or nil.
	orderIdtoDel := ""
	// here we will prepare the url with parameters to pass to our request
	url := "/order?" + "id=" + orderIdtoDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	OrderDelete(w, request)

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

// Test3 check sql error foreign key constraint violation(1451),show 500.
func TestOrderSqlErrorCheck(t *testing.T) {

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	log.Println("mockErr: ", err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I  expect a Delete Query execution and for that :
	// dbmock.ExpectExec("DELETE FROM orders").WillReturnResult(sqlmock.NewResult(1, 1))
	expectedErr := mysql.MySQLError{Number: 1451, Message: "Foreign key constraint violation."}
	dbmock.ExpectExec("DELETE FROM orders").
		WillReturnError(&expectedErr)

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// set order id empty or nil.
	orderIdtoDel := "1"
	// here we will prepare the url with parameters to pass to our request
	url := "/order?" + "id=" + orderIdtoDel

	// Create a mock Request
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	OrderDelete(w, request)

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
