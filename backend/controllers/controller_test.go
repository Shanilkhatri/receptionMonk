package controllers

import "net/http"

type MockHelper struct {
	MockReturnUserDetailsResult error
}

func (m MockHelper) ReturnUserDetails(r *http.Request, user interface{}) error {
	return m.MockReturnUserDetailsResult
}
