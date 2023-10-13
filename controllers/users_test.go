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
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

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
// -> user updating his own record
// -> we are updating 5 things (name,pass,key,recovCode,dob)
// -> expecting a 200
func TestUserPostWithCorrectData(t *testing.T) {
	// data to be posted
	jsonData := map[string]interface{}{
		"id":                    1,
		"name":                  "shaanil", // changed name
		"email":                 "user@example.com",
		"passwordHash":          "1234",       // changing pass
		"twoFactorKey":          "55",         // changing key
		"twoFactorRecoveryCode": "59898",      // changing code
		"dob":                   "2023-10-06", // changed dob
		"accountType":           "user",
		"companyId":             2,
		"status":                "active",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as user
	var userdetails utility.UserDetails
	userdetails.ID = 1
	userdetails.AccountType = "user"
	userdetails.CompanyID = 2
	userdetails.DOB = "2023-10-05" // dob at DB/cache
	userdetails.Name = "hguhduhs"  // name at DB/cache
	userdetails.Email = "user@example.com"
	userdetails.TwoFactorKey = "iuriouf08959374rvseuyyrv94w857yesiufhu" //key at DB/cache
	userdetails.TwoFactorRecoveryCode = "fc78"                          // twoFactRecCode at DB/cache
	userdetails.PasswordHash = "dihfw94534yrehu8y348vy3uy84728"         // passHash at DB/cache
	userdetails.Status = "active"

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	// log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("UPDATE `authentication`").WillReturnResult(sqlmock.NewResult(1, 1))

	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostUser(w, request)
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

// TEST #2 postUser with correct data and struct
// -> user updating another user
// -> we are updating 5 things here also (name,pass,key,recovCode,dob)
// -> won't be mocking the DB as we expect to catch error before DB ops
// -> expecting a 403
func TestUserPostWithUserUpdatingAnotherUser(t *testing.T) {
	// data to be posted
	jsonData := map[string]interface{}{
		"id":                    2,         // trying to update id #2 data
		"name":                  "shaanil", // changed name
		"email":                 "user@example.com",
		"passwordHash":          "1234",       // changing pass
		"twoFactorKey":          "55",         // changing key
		"twoFactorRecoveryCode": "59898",      // changing code
		"dob":                   "2023-10-06", // changed dob
		"accountType":           "user",
		"companyId":             2,
		"status":                "active",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as user
	var userdetails utility.UserDetails
	userdetails.ID = 1 // the user's actual id
	userdetails.AccountType = "user"
	userdetails.CompanyID = 2
	userdetails.DOB = "2023-10-05" // dob at DB/cache
	userdetails.Name = "hguhduhs"  // name at DB/cache
	userdetails.Email = "user@example.com"
	userdetails.TwoFactorKey = "iuriouf08959374rvseuyyrv94w857yesiufhu" //key at DB/cache
	userdetails.TwoFactorRecoveryCode = "fc78"                          // twoFactRecCode at DB/cache
	userdetails.PasswordHash = "dihfw94534yrehu8y348vy3uy84728"         // passHash at DB/cache
	userdetails.Status = "active"

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostUser(w, request)

	// we can even read body using io package and test even specific messages
	// for now I have skipped it, might be added in future

	if w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Result().StatusCode)
	}
}

// TEST #3 postUser with correct data and struct
// -> owner updating another user of the same company
// -> we are updating 5 things here also (name,pass,key,recovCode,dob)
// -> will be mocking DB ops as we are expecting successfull updation.
// -> expecting a 200

func TestUserPostWithOwnerUpdUserOfSameComp(t *testing.T) {
	// data to be posted
	jsonData := map[string]interface{}{
		"id":                    1,
		"name":                  "shaanil", // changed name
		"email":                 "user@example.com",
		"passwordHash":          "1234",       // changing pass
		"twoFactorKey":          "55",         // changing key
		"twoFactorRecoveryCode": "59898",      // changing code
		"dob":                   "2023-10-06", // changed dob
		"accountType":           "user",
		"companyId":             2,
		"status":                "active",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as user
	var userdetails utility.UserDetails
	userdetails.ID = 3                // diff user
	userdetails.AccountType = "owner" // type set to owner
	userdetails.CompanyID = 2         // belong to same company
	userdetails.DOB = "2023-10-05"    // dob at DB/cache
	userdetails.Name = "hguhduhs"     // name at DB/cache
	userdetails.Email = "user@example.com"
	userdetails.TwoFactorKey = "iuriouf08959374rvseuyyrv94w857yesiufhu" //key at DB/cache
	userdetails.TwoFactorRecoveryCode = "fc78"                          // twoFactRecCode at DB/cache
	userdetails.PasswordHash = "dihfw94534yrehu8y348vy3uy84728"         // passHash at DB/cache
	userdetails.Status = "active"

	// open Mock DB connection
	mockDB, dbmock, err := sqlmock.New()
	// log.Println(err)
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	// mocking the user detail I expect from the Db op
	expectedUser := models.Users{
		ID:                    1,
		AccountType:           "user",
		CompanyID:             2,
		DOB:                   "2023-10-05",
		Name:                  "hguhduhs",
		Email:                 "user@example.com",
		TwoFactorKey:          "iuriouf08959374rw857yesiufhu",
		TwoFactorRecoveryCode: "fjjdjfn",
		PasswordHash:          "dihfw94534yrehu8yuy84728",
		Status:                "active",
	}

	// make expected user a row that will be returned
	rows := sqlmock.NewRows([]string{"id", "name", "email", "passwordHash", "twoFactorKey", "twoFactorRecoveryCode", "dob", "accountType", "companyId", "status"}).AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email, expectedUser.PasswordHash, expectedUser.TwoFactorKey, expectedUser.TwoFactorRecoveryCode, expectedUser.DOB, expectedUser.AccountType, expectedUser.CompanyID, expectedUser.Status)

	// I am expecting a db op before final updation to get user
	dbmock.ExpectQuery("SELECT \\* FROM authentication WHERE id = ?").WillReturnRows(rows)
	// I have used mustBegin thats why I am using Expect begin
	dbmock.ExpectBegin()
	// I also expect an Insert Query execution and for that :
	dbmock.ExpectExec("UPDATE `authentication`").WillReturnResult(sqlmock.NewResult(1, 1))

	// expecting a commit to as this is correct info
	dbmock.ExpectCommit()

	// Binding the DB Cursor to correct utility.Db
	utility.Db = sqlxDB

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostUser(w, request)
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

// TEST #4 postUser with correct data and struct
// -> owner updating another user of different company
// -> we are updating 5 things here also (name,pass,key,recovCode,dob)
// -> will not be mocking DB ops as we are expecting an error before that.
// -> expecting a 403

func TestUserPostWithOwnerUpdUserOfDiffComp(t *testing.T) {
	// data to be posted
	jsonData := map[string]interface{}{
		"id":                    1,
		"name":                  "shaanil", // changed name
		"email":                 "user@example.com",
		"passwordHash":          "1234",       // changing pass
		"twoFactorKey":          "55",         // changing key
		"twoFactorRecoveryCode": "59898",      // changing code
		"dob":                   "2023-10-06", // changed dob
		"accountType":           "user",
		"companyId":             3, // company id is diff from that of owner
		"status":                "active",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as owner
	var userdetails utility.UserDetails
	userdetails.ID = 3                // diff user
	userdetails.AccountType = "owner" // type set to owner
	userdetails.CompanyID = 2         // belong to diff company company
	userdetails.DOB = "2023-10-05"    // dob at DB/cache
	userdetails.Name = "hguhduhs"     // name at DB/cache
	userdetails.Email = "user@example.com"
	userdetails.TwoFactorKey = "iuriouf08959374rvseuyyrv94w857yesiufhu" //key at DB/cache
	userdetails.TwoFactorRecoveryCode = "fc78"                          // twoFactRecCode at DB/cache
	userdetails.PasswordHash = "dihfw94534yrehu8y348vy3uy84728"         // passHash at DB/cache
	userdetails.Status = "active"

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostUser(w, request)

	// we can even read body using io package and test even specific messages
	// for now I have skipped it, might be added in future

	if w.Result().StatusCode != 403 {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, w.Result().StatusCode)
	}
}

// TEST #5 postUser with incorrect data format (bad email string)
// -> owner updating user details with a faulty email
// -> we are updating 5 things here also (name,pass,key,recovCode,dob) but will get caught only at email
// and to proove that we'll decode the body and also read the response.Message to make sure we got stuck at
// faulty mail
// -> will not be mocking DB ops as we are expecting an error before that.
// -> expecting a 400

func TestUserPostWithOwnerUpdUserWithFaultyDate(t *testing.T) {
	// data to be posted
	jsonData := map[string]interface{}{
		"id":                    1,
		"name":                  "shaanil",         // changed name
		"email":                 "user@examplecom", // ----------> faulty (missing .)
		"passwordHash":          "1234",            // changing pass
		"twoFactorKey":          "55",              // changing key
		"twoFactorRecoveryCode": "59898",           // changing code
		"dob":                   "2023-10-06",      // changed dob
		"accountType":           "user",
		"companyId":             2,
		"status":                "active",
	}
	// Marshal the data into JSON format
	requestBody, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// mocking token payload with userDetails as owner
	var userdetails utility.UserDetails
	userdetails.ID = 3                // diff user
	userdetails.AccountType = "owner" // type set to owner
	userdetails.CompanyID = 2
	userdetails.DOB = "2023-10-05" // dob at DB/cache
	userdetails.Name = "hguhduhs"  // name at DB/cache
	userdetails.Email = "user@example.com"
	userdetails.TwoFactorKey = "iuriouf08959374rvseuyyrv94w857yesiufhu" //key at DB/cache
	userdetails.TwoFactorRecoveryCode = "fc78"                          // twoFactRecCode at DB/cache
	userdetails.PasswordHash = "dihfw94534yrehu8y348vy3uy84728"         // passHash at DB/cache
	userdetails.Status = "active"

	// Mocking the utility functions that are used there
	Utility = MockHelper{
		// MockStrictParseDataFromJsonResult:      nil,
		// MockSessionGetResult:                   "owner", //setting session won't be neccessary here
		MockCheckTokenPayloadAndReturnUserBool:    true,
		MockCheckTokenPayloadAndReturnUserDetails: userdetails,
	}
	// Create a mock Request
	request := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	// Call your function with the mocks
	PostUser(w, request)

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

	if w.Result().StatusCode != 400 && data.Message != "Please enter valid email address" {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Result().StatusCode)
	}
}
