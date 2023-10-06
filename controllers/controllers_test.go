package controllers

import (
	"net/http/httptest"
	"reakgo/models"
	"testing"
)

func TestPostExtension(t *testing.T) {

	// Create a mock HTTP request
	request := httptest.NewRequest("GET", "/example", nil)

	// Create a mock HTTP response recorder
	recorder := httptest.NewRecorder()
	testCases := []struct {
		userDetails models.Users
		expected    bool
	}{
		// test case 1 for valid user .
		{models.Users{Id: 1, CompanyId: 2, AccountType: "owner"}, true},

		// test case  2 for invalid user.
		{models.Users{Id: 0, CompanyId: 2, AccountType: "owner"}, false},

		// test case 3 for invalid user.
		{models.Users{Id: 1, CompanyId: 0, AccountType: "owner"}, false},

		// test case  4 for invalid user.
		{models.Users{Id: 1, CompanyId: 2, AccountType: "unknown"}, false},

		// test case 5 for valid user.
		{models.Users{Id: 1, CompanyId: 2, AccountType: "user"}, true},

		// test case  6 for invalid user.
		{models.Users{Id: -1, CompanyId: 2, AccountType: "owner"}, false},

		// test case 7 for invalid user.
		{models.Users{Id: 1, CompanyId: -2, AccountType: "user"}, false},

		// test case 8 for invalid user.
		{models.Users{Id: 1, CompanyId: 2, AccountType: "admin"}, false},

		// test case 9 for valid user .
		{models.Users{Id: 1, CompanyId: 2, AccountType: "user"}, true},
	}

	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			//main function call for testing.
			result := PostExtension(recorder, request, testCase.userDetails)
			if result != testCase.expected {
				t.Errorf("For input %+v, expected %v but got %v", testCase.userDetails, testCase.expected, result)
			}
		})
	}
}
