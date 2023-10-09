package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
)

func AddForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			// Logic
			name := r.FormValue("name")
			address := r.FormValue("address")

			log.Println(name, address)
			models.FormAddView{}.Add(name, address)
		}
	}
	Utility.RenderTemplate(w, r, "addForm", nil)
}

func ViewForm(w http.ResponseWriter, r *http.Request) {
	result, err := models.FormAddView{}.View()
	if err != nil {
		log.Println(err)
	}
	Utility.RenderTemplate(w, r, "viewForm", result)
}
