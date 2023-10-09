package controllers

import (
	"net/http"
)

func BaseIndex(w http.ResponseWriter, r *http.Request) {
	name := []string{"Test1", "Test2"}
	Utility.RenderTemplate(w, r, "index", name)
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	Utility.RenderTemplate(w, r, "dashboard", nil)
}
