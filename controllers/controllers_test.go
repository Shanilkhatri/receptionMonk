// package controllers

// import (
// 	"net/http/httptest"
// 	"reakgo/models"
// 	"testing"
// )

// func TestPostExtension(t *testing.T) {

// 	// Create a mock HTTP request
// 	request := httptest.NewRequest("GET", "/example", nil)

// 	// Create a mock HTTP response recorder
// 	recorder := httptest.NewRecorder()
// 	testCases := []struct {
// 		userDetails models.Users
// 		expected    bool
// 	}{
// 		// test case 1 for valid user .
// 		{models.Users{Id: 1, CompanyId: 2, AccountType: "owner"}, true},

// 		// test case  2 for invalid user.
// 		{models.Users{Id: 0, CompanyId: 2, AccountType: "owner"}, false},

// 		// test case 3 for invalid user.
// 		{models.Users{Id: 1, CompanyId: 0, AccountType: "owner"}, false},

// 		// test case  4 for invalid user.
// 		{models.Users{Id: 1, CompanyId: 2, AccountType: "unknown"}, false},

// 		// test case 5 for valid user.
// 		{models.Users{Id: 1, CompanyId: 2, AccountType: "user"}, true},

// 		// test case  6 for invalid user.
// 		{models.Users{Id: -1, CompanyId: 2, AccountType: "owner"}, false},

// 		// test case 7 for invalid user.
// 		{models.Users{Id: 1, CompanyId: -2, AccountType: "user"}, false},

// 		// test case 8 for invalid user.
// 		{models.Users{Id: 1, CompanyId: 2, AccountType: "admin"}, false},

// 		// test case 9 for valid user .
// 		{models.Users{Id: 1, CompanyId: 2, AccountType: "user"}, true},
// 	}

// 	for _, testCase := range testCases {
// 		t.Run("", func(t *testing.T) {
// 			//main function call for testing.
// 			result := PostExtension(recorder, request)
// 			if result != testCase.expected {
// 				t.Errorf("For input %+v, expected %v but got %v", testCase.userDetails, testCase.expected, result)
// 			}
// 		})
// 	}
// }

// // func createRequest(userPayload models.Users, extenPayload models.Extensions) (*http.Request, error) {
// // 	userJSON, err := json.Marshal(userPayload)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	extenJSON, err := json.Marshal(extenPayload)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	req := httptest.NewRequest("POST", "/example", bytes.NewReader(extenJSON))
// // 	req.Header.Add("tokenPayload", string(userJSON))

// // 	authorization := " Bearer YourNewTokenValue6789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmno"
// // 	req.Header.Add("Authorization", authorization)

// // 	return req, nil

// // }

// // func TestPostExtensionSuccess(t *testing.T) {

// // 	userPayload := createUserPayload(-1, "name", "email", "passwordHash", "twofactorkey", " twoFactorRecoveryCode", "dob", "owner", 1, "active", "etryerguyerfdjfdffjioetyutrueurtrur")

// // 	extenPayload := createExtenPayload(1, "extension", 1, 1, "sipserver", "sipusename", "sipPassword", "sipport")

// // 	req, err := createRequest(userPayload, extenPayload)
// // 	if err != nil {
// // 		t.Fatalf("createRequest error: %v", err)
// // 	}

// // 	log.Println("req", req)

// // 	recorder := httptest.NewRecorder()
// // 	PostExtension(recorder, req)
// // 	assert.Equal(t, http.StatusOK, recorder.Code)

// // }

// // func createUserPayload(id int, name string, email string, passwordHash string, twoFactorKey string, twoFactorRecoveryCode string, dob string, accountType string, compantId int, status string, token string) models.Users {
// // 	return models.Users{
// // 		Id:                    id,
// // 		Name:                  name,
// // 		Email:                 email,
// // 		PasswordHash:          passwordHash,
// // 		TwoFactorKey:          twoFactorKey,
// // 		TwoFactorRecoveryCode: twoFactorRecoveryCode,
// // 		DOB:                   dob,
// // 		AccountType:           accountType,
// // 		CompanyId:             compantId,
// // 		Status:                status,
// // 		Token:                 token,
// // 	}
// // }

// // func createExtenPayload(id int, extension string, userId int, department int, sipserver string, sipusername string, sippassword string, sipport string) models.Extensions {
// // 	return models.Extensions{
// // 		Id:          id,
// // 		Extension:   extension,
// // 		UserId:      userId,
// // 		Department:  department,
// // 		SipServer:   sipserver,
// // 		SipUserName: sipusername,
// // 		SipPassword: sippassword,
// // 		SipPort:     sipport,
// // 	}
// // }

// // func TestPostExtensionErrorCases(t *testing.T) {
// // 	invalidUserPayload := createUserPayload(1, "name", "email", "passwordHash", "twofactorkey", " twoFactorRecoveryCode", "dob", "owner", 1, "active", "etryerguyerfdjfdffjioetyutrueurtrur")
// // 	invalidExtenPayload := createExtenPayload(1, "extension", 1, 1, "sipserver", "sipusename", "sipPassword", "sipport")

// // 	invalidReq, err := createRequest(invalidUserPayload, invalidExtenPayload)
// // 	if err != nil {
// // 		t.Fatalf("createRequest error: %v", err)
// // 	}

// // 	invalidRecorder := httptest.NewRecorder()

// // 	PostExtension(invalidRecorder, invalidReq)

// // 	assert.Equal(t, http.StatusInternalServerError, invalidRecorder.Code)

// // 	expectedErrorMessage := "Validation failed"
// // 	assert.Contains(t, invalidRecorder.Body.String(), expectedErrorMessage)

// // }

// // // func createUserPayload(id int, accountType string, compantId int) models.Users {
// // // 	return models.Users{
// // // 		Id:          id,
// // // 		AccountType: accountType,
// // // 		CompanyId:   compantId,
// // // 	}
// // // }
