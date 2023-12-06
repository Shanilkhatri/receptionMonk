package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reakgo/utility"
	"reflect"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type MockHelper struct {
	MockReturnUserDetailsResult                error
	MockGenerateRandomStringStrResult          string
	MockGenerateRandomStringErrorResult        error
	MockSessionGetResult                       interface{}
	MockViewFlashResult                        interface{}
	MockParseDataFromPostRequestToMapResult    map[string]interface{}
	MockParseDataFromPostRequestToMapErrResult error
	MockParseDataFromJsonToMapResult           map[string]interface{}
	MockParseDataFromJsonToMapErrResult        error
	MockStrictParseDataFromJsonResult          error
	MockStrictParseDataFromPostRequestResult   error
	MockStringInArray                          bool
	MockCheckTokenPayloadAndReturnUserBool     bool
	MockCheckTokenPayloadAndReturnUserDetails  utility.UserDetails
	MockGetSqlErrorString                      string
	MockCheckDateFormat                        bool
	MockCheckEmailFormat                       bool
	MockDeleteSessionValues                    bool
}

// CheckDateFormat implements utility.Helper.
func (m MockHelper) CheckDateFormat(dateString string) bool {
	return m.MockCheckDateFormat
}

// CheckEmailFormat implements utility.Helper.
func (m MockHelper) CheckEmailFormat(emailString string) bool {
	return m.MockCheckEmailFormat
}

// CheckSqlError implements utility.Helper.
func (MockHelper) CheckSqlError(err error, errString string) (bool, string) {
	panic("unimplemented")
}

// CopyFieldsBetweenDiffStructType implements utility.Helper.
func (m MockHelper) CopyFieldsBetweenDiffStructType(src interface{}, dest interface{}) bool {
	srcValue := reflect.ValueOf(src)
	destValue := reflect.ValueOf(dest).Elem() // Use Elem to get the underlying struct Value.

	if srcValue.Kind() != reflect.Struct || destValue.Kind() != reflect.Struct {
		fmt.Println("Both src and dest should be structs")
		return false
	}

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		destField := destValue.Field(i)

		// Check if the field in dest is assignable from the field in src
		if destField.Type().AssignableTo(srcField.Type()) {
			destField.Set(srcField)
		}
	}
	return true
}

// DeleteSessionValues implements utility.Helper.
func (m MockHelper) DeleteSessionValues(w http.ResponseWriter, r *http.Request, KeyName string) bool {
	return m.MockDeleteSessionValues
}

// FillEmptyFieldsForPostUpdate implements utility.Helper.
func (MockHelper) FillEmptyFieldsForPostUpdate(src interface{}, dest interface{}) bool {
	srcValue := reflect.ValueOf(src)
	destValue := reflect.ValueOf(dest).Elem() // Use Elem to get the underlying struct Value.

	if srcValue.Kind() != reflect.Struct || destValue.Kind() != reflect.Struct {
		log.Println("Both src and dest should be structs")
		return false
	}

	if srcValue.Type() != destValue.Type() {
		log.Println("src and dest should have the same struct type")
		return false
	}

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		destField := destValue.Field(i)
		if destField.IsZero() {
			// If empty, fill it with the value from src
			destField.Set(srcField)
		}

	}
	return true
}

// GetErrorMessage implements utility.Helper.
func (MockHelper) GetErrorMessage(currentFilePath string, lineNumbers int, errorMessage error) bool {
	panic("unimplemented")
}

// GetImageTypeExtension implements utility.Helper.
func (MockHelper) GetImageTypeExtension(Filename string, whatToBeTrim string, dotInclude bool) string {
	panic("unimplemented")
}

// GetSqlErrorString implements utility.Helper.
func (m MockHelper) GetSqlErrorString(err error) string {
	return m.MockGetSqlErrorString
}

// Logger implements utility.Helper.
func (MockHelper) Logger(errObject error, flag bool) {
	panic("unimplemented")
}

// NewPasswordHash implements utility.Helper.
func (m MockHelper) NewPasswordHash(NewPassword string) (string, error) {
	//NewPassword Change bcrypt code
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(NewPassword), 10)
	//modify NewPassword
	NewPassword = string(newPasswordHash)
	if err != nil || NewPassword == "" {
		m.Logger(err, true)
	} else {
		return NewPassword, err
	}
	return "", err
}

// OpenLogFile implements utility.Helper.
func (m MockHelper) OpenLogFile() *os.File {
	panic("unimplemented")
}

// SaltPlainPassWord implements utility.Helper.
func (m MockHelper) SaltPlainPassWord(passW string) (string, error) {
	// making hash of pass #1
	hashedPassW, err := m.NewPasswordHash(passW)
	if err != nil {
		return "", err
	}
	// mixing salt with hashed pass
	pswdConcatWithSalt := hashedPassW + os.Getenv("CONS_SALT")

	// making hash of (salted+hashed) pass #2
	hashedPassW, err = m.NewPasswordHash(pswdConcatWithSalt)
	if err != nil {
		return "", err
	}
	return hashedPassW, nil
}

// SendEmail implements utility.Helper.
func (MockHelper) SendEmail(to []string, template string, data map[string]interface{}) (int, bool, error) {
	return 0, true, nil
}

// SendEmailSMTP implements utility.Helper.
func (MockHelper) SendEmailSMTP(to []string, subject string, body string) (bool, error) {
	return true, nil
}

// StrToInt implements utility.Helper.
func (m MockHelper) StrToInt(num string) int {
	if num != "" {
		intNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
		}
		return intNum
	}
	return 0
}

// StrToInt64 implements utility.Helper.
func (m MockHelper) StrToInt64(str string) (int64, error) {
	strint64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return int64(0), err
	}
	return strint64, err
	// panic("unimplemented")
}

// type Session struct {
// 	Key   string
// 	Value interface{}
// }

func (m MockHelper) ReturnUserDetails(r *http.Request, user interface{}) error {
	return m.MockReturnUserDetailsResult
}

func (m MockHelper) AddFlash(flavour string, message string, w http.ResponseWriter, r *http.Request) {

}

func (m MockHelper) GenerateRandomString(n int) (string, error) {
	return m.MockGenerateRandomStringStrResult, m.MockGenerateRandomStringErrorResult
}

func (m MockHelper) RedirectTo(w http.ResponseWriter, r *http.Request, path string) {

}

func (m MockHelper) SessionGet(r *http.Request, key string) interface{} {
	return m.MockSessionGetResult
}

func (m MockHelper) SessionSet(w http.ResponseWriter, r *http.Request, data utility.Session) {

}

func (m MockHelper) ViewFlash(w http.ResponseWriter, r *http.Request) interface{} {
	return m.MockViewFlashResult
}

func (m MockHelper) RenderTemplate(w http.ResponseWriter, r *http.Request, template string, data interface{}) {

}

func (m MockHelper) ParseDataFromPostRequestToMap(r *http.Request) (map[string]interface{}, error) {
	return m.MockParseDataFromPostRequestToMapResult, m.MockParseDataFromPostRequestToMapErrResult
}

func (m MockHelper) ParseDataFromJsonToMap(r *http.Request) (map[string]interface{}, error) {
	return m.MockParseDataFromJsonToMapResult, m.MockParseDataFromJsonToMapErrResult
}

func (m MockHelper) StrictParseDataFromJson(r *http.Request, structure interface{}) error {
	err := json.NewDecoder(r.Body).Decode(structure)
	if err != nil {
		return err
	}

	return err
}

func (m MockHelper) StrictParseDataFromPostRequest(r *http.Request, structure interface{}) error {
	return m.MockStrictParseDataFromPostRequestResult
}

func (m MockHelper) RenderJsonResponse(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	jsonresponce, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// we take the statusCode as an arguement and send it as a http response
	log.Println("statusCode: ", statusCode)
	switch statusCode {
	case 403:
		w.WriteHeader(http.StatusForbidden)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 500:
		w.WriteHeader(http.StatusInternalServerError)
	case 200:
		w.WriteHeader(http.StatusOK)
	}
	w.Write([]byte(jsonresponce))
}

func (m MockHelper) RenderTemplateData(w http.ResponseWriter, r *http.Request, template string, data interface{}) {

}

func (m MockHelper) StringInArray(target string, arr []string) bool {
	return m.MockStringInArray
}

func (m MockHelper) CheckTokenPayloadAndReturnUser(r *http.Request) (bool, utility.UserDetails) {

	return m.MockCheckTokenPayloadAndReturnUserBool, m.MockCheckTokenPayloadAndReturnUserDetails
}
