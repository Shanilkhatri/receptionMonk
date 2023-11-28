package controllers

import (
	"net/http"
	"os"
	"reakgo/utility"
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
	// MockStrictParseDataFromJsonResult          error
	MockStrictParseDataFromPostRequestResult  error
	MockStringInArray                         bool
	MockCheckTokenPayloadAndReturnUserBool    bool
	MockCheckTokenPayloadAndReturnUserDetails utility.UserDetails
}

// CheckDateFormat implements utility.Helper.
func (MockHelper) CheckDateFormat(dateString string) bool {
	panic("unimplemented")
}

// CheckEmailFormat implements utility.Helper.
func (MockHelper) CheckEmailFormat(emailString string) bool {
	panic("unimplemented")
}

// CheckSqlError implements utility.Helper.
func (MockHelper) CheckSqlError(err error, errString string) (bool, string) {
	panic("unimplemented")
}

// CopyFieldsBetweenDiffStructType implements utility.Helper.
func (MockHelper) CopyFieldsBetweenDiffStructType(src interface{}, dest interface{}) bool {
	panic("unimplemented")
}

// DeleteSessionValues implements utility.Helper.
func (MockHelper) DeleteSessionValues(w http.ResponseWriter, r *http.Request, KeyName string) bool {
	panic("unimplemented")
}

// FillEmptyFieldsForPostUpdate implements utility.Helper.
func (MockHelper) FillEmptyFieldsForPostUpdate(src interface{}, dest interface{}) bool {
	panic("unimplemented")
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
func (MockHelper) GetSqlErrorString(err error) string {
	panic("unimplemented")
}

// Logger implements utility.Helper.
func (MockHelper) Logger(errObject error) {
	panic("unimplemented")
}

// NewPasswordHash implements utility.Helper.
func (MockHelper) NewPasswordHash(NewPassword string) (string, error) {
	panic("unimplemented")
}

// OpenLogFile implements utility.Helper.
func (MockHelper) OpenLogFile() *os.File {
	panic("unimplemented")
}

// SaltPlainPassWord implements utility.Helper.
func (MockHelper) SaltPlainPassWord(passW string) (string, error) {
	panic("unimplemented")
}

// SendEmail implements utility.Helper.
func (MockHelper) SendEmail(to []string, template string, data map[string]interface{}) bool {
	panic("unimplemented")
}

// SendEmailSMTP implements utility.Helper.
func (MockHelper) SendEmailSMTP(to []string, subject string, body string) bool {
	panic("unimplemented")
}

// StrToInt implements utility.Helper.
func (MockHelper) StrToInt(num string) int {
	panic("unimplemented")
}

// StrToInt64 implements utility.Helper.
func (MockHelper) StrToInt64(str string) (int64, error) {
	panic("unimplemented")
}

// StrictParseDataFromJson implements utility.Helper.
func (MockHelper) StrictParseDataFromJson(r *http.Request, structure interface{}) error {
	panic("unimplemented")
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

// func (m MockHelper) StrictParseDataFromJson(r *http.Request, structure interface{}) error {
// 	return m.MockStrictParseDataFromJsonResult
// }

func (m MockHelper) StrictParseDataFromPostRequest(r *http.Request, structure interface{}) error {
	return m.MockStrictParseDataFromPostRequestResult
}

func (m MockHelper) RenderJsonResponse(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {

}

func (m MockHelper) RenderTemplateData(w http.ResponseWriter, r *http.Request, template string, data interface{}) {

}

func (m MockHelper) StringInArray(target string, arr []string) bool {
	return m.MockStringInArray
}

func (m MockHelper) CheckTokenPayloadAndReturnUser(r *http.Request) (bool, utility.UserDetails) {

	return m.MockCheckTokenPayloadAndReturnUserBool, m.MockCheckTokenPayloadAndReturnUserDetails
}
