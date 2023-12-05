package controllers

import (
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
	"github.com/jmoiron/sqlx"
)

// TEST #1 - GetOrderDetails with correct Struct and also correct Data,show 200.
func TestGetOrderDetailsSuccess(t *testing.T) {

	queryParams := map[string]string{
		"id":        "1",
		"companyId": "1",
	}

	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 1

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()

	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	expectedOrderDetails := models.OrderDetails{
		Id:              1,
		OrderId:         3,
		PhoneNumber:     "8768990895",
		SipServer:       "sipServer",
		SipUsername:     "SipUser@name",
		SipPassword:     "string@123#$",
		SipPort:         "5678",
		IsIvrEnabled:    true,
		IvrFlow:         "string",
		MaxAllowedUsers: 0,
		MaxAllowedDepts: 0,
	}

	// make expected user a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "orderId", "phoneNumber", "sipServer", "sipUsername", "sipPassword", "sipPort", "isIvrEnabled", "ivrFlow", "maxAllowedUsers", "maxAllowedDepts"}).AddRow(expectedOrderDetails.Id, expectedOrderDetails.OrderId, expectedOrderDetails.PhoneNumber, expectedOrderDetails.SipServer, expectedOrderDetails.SipUsername, expectedOrderDetails.SipPassword, expectedOrderDetails.SipPort, expectedOrderDetails.IsIvrEnabled, expectedOrderDetails.IvrFlow, expectedOrderDetails.MaxAllowedUsers, expectedOrderDetails.MaxAllowedDepts)

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM OrderDetails ").WillReturnRows(rows)
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
	// here we will prepare the url with parameters to pass to our request.
	url := "/orderdetails?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	OrderDetailsGet(w, request)

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

	if data.Message != "Returns all matching order details." && w.Result().StatusCode != 200 {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
	}
}

// TEST #2 - GetOrderDetails if accountType not equal to owner,show 403.
func TestGetOrderDetailsAccountTypeIsUser(t *testing.T) {

	queryParams := map[string]string{
		"id":        "1",
		"companyId": "1",
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
	// here we will prepare the url with parameters to pass to our request.
	url := "/orderdetails?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	OrderDetailsGet(w, request)

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

// TEST #3 - GetOrderDetails Sql return no result found,show 500.
func TestGetOrderDetailsReturnSqlErrors(t *testing.T) {

	queryParams := map[string]string{
		"id":        "1",
		"companyId": "1",
	}

	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "owner"
	userdetails.CompanyID = 1

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()

	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I also expect a GET Query execution and for that :
	dbmock.ExpectQuery("SELECT \\* FROM OrderDetails").WillReturnError(sql.ErrNoRows) // Simulate a "no rows found" error.

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
	// here we will prepare the url with parameters to pass to our request.
	url := "/orderdetails?" + buildQuery(queryParams)
	// Create a mock Request
	request := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	// Call your function with the mocks
	OrderDetailsGet(w, request)

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
