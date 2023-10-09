package controllers

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http/httptest"
	"testing"

	"reakgo/models"

	"github.com/stretchr/testify/mock"
)

// Standard Definitions

// Mock DB Function
type MockExtensionPoster struct {
	mock.Mock
}

func (m *MockExtensionPoster) PostExtension(ext models.Extensions, tx *sql.Tx) (bool, error) {
	args := m.Called(ext, tx)
	return args.Bool(0), args.Error(1)
}

// Tests the add extension function by sending correct parameters to see if we get proper response
func TestExtensionAddWithIncorrectData(t *testing.T) {
	// Create a mock HTTP request
	jsonData := `{
	"id": 1,
  	"first_name": "Jeanette",
  	"last_name": "Penddreth",
  	"email": "jpenddreth0@census.gov",
  	"gender": "Female",
  	"ip_address": "26.58.193.2"
}`
	r := httptest.NewRequest("POST", "/extension/add", bytes.NewBuffer([]byte(jsonData)))

	// Create a mock HTTP response writer
	w := httptest.NewRecorder()

	_ = PostExtension(w, r)

	if w.Result().Status != fmt.Sprint(400) {
		t.Errorf("Expected controller to return 400, got %s", w.Result().Status)
	}
}

func TestExtensionAddWithCorrectData(t *testing.T) {
	// Create a mock HTTP request
	jsonData := `{
	"id": 1,
  	"first_name": "Jeanette",
  	"last_name": "Penddreth",
  	"email": "jpenddreth0@census.gov",
  	"gender": "Female",
  	"ip_address": "26.58.193.2"
}`

	r := httptest.NewRequest("POST", "/extension/add", bytes.NewBuffer([]byte(jsonData)))

	// Create a mock HTTP response writer
	w := httptest.NewRecorder()

	_ = PostExtension(w, r)

	if w.Result().Status != fmt.Sprint(200) {
		t.Errorf("Expected controller to return 200, got %s", w.Result().Status)
	}
}
