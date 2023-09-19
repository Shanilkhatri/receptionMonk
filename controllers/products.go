package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"reakgo/models"
)

func ProductJsonDecoder(r *http.Request) (models.Products, error) {
	var productStruct models.Products
	err := json.NewDecoder(r.Body).Decode(&productStruct)
	return productStruct, err
}

func PutProducts(w http.ResponseWriter, r *http.Request) bool {
	// response := utility.AjaxResponse{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	productStruct, err := ProductJsonDecoder(r)
	if err != nil {
		// response.Message = err.Error()
		// utility.RenderTemplate(w, r, "", response)
		return true
	}
	// debug log
	log.Println("productStruct", productStruct)
	// if values are not valid return
	if productStruct.Name == "" || productStruct.Price == 0 || productStruct.PlanValidity == 0 || productStruct.Status == "" {
		// response.Message = "Please fill all the required fields with valid values."
		// utility.RenderTemplate(w, r, "", response)
		return true
	}
	// else insert product details
	boolVal, err := Db.products.PutProducts(productStruct)

	if boolVal || err != nil {
		// utility.Logger(err)
		// response.Message = "Unable to save data, The system encountered an issue while attempting to store the data."
		// utility.RenderTemplate(w, r, "", response)
		return true
	}
	// response.Status = "200"
	// response.Message = "Product added successfully."
	// utility.RenderTemplate(w, r, "", nil)
	return false
}

// func PostProducts(w http.ResponseWriter, r *http.Request) bool {
// 	exec(restartCommand, (error, stdout, stderr) => {
// 		if (error) {
// 			console.error(`Error restarting Node-RED: ${error}`);
// 			return;
// 		}
// 		console.log(`Node-RED restarted: ${stdout}`);
// 	});
// }
