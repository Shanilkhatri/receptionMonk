package controllers

import (
	// Standard Library Packages
	"bytes"
	"fmt"
	"log"
	"net/http/httptest"
	"reakgo/utility"
	"testing"

	// Third Party Packages
	"github.com/DATA-DOG/go-sqlmock"
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

	Utility = MockHelper{
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
