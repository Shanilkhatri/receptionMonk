package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reakgo/models"
	"reakgo/utility"
)

func PutKycDetails(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var userStruct models.KycDetails
	err := utility.StrictParseDataFromJson(r, &userStruct)
	if err != nil {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	isok, userDetailsType := utility.CheckTokenPayloadAndReturnUser(r)
	if !isok {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	userStruct.UserId = userDetailsType.ID
	if userStruct.UserId > 0 && userStruct.DocPicName != "" && userStruct.DocName != "" {
		log.Println("userStruct: ", userStruct)
		tx := utility.Db.MustBegin()
		inserted := models.KycDetails{}.Putkyc(userStruct, tx)
		if inserted {
			err = tx.Commit()
			if err != nil {
				log.Println(err)
				tx.Rollback()
				response.Status = "400"
				response.Message = "Unable to update kyc at the moment! Please try again."
				utility.RenderJsonResponse(w, r, response, 400)
				return
			}
			response.Status = "200"
			response.Message = "Document upload successfully"
			utility.RenderJsonResponse(w, r, response, 200)
			return
		} else {
			tx.Rollback()
			response.Status = "400"
			response.Message = "Unable to update kyc at the moment! Please try again."
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
	} else {
		response.Status = "400"
		response.Message = "Please provide all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
}
func PostKycDetails(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var userStruct models.KycDetails
	err := utility.StrictParseDataFromJson(r, &userStruct)
	if err != nil {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	isok, userDetailsType := utility.CheckTokenPayloadAndReturnUser(r)
	if !isok {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	userStruct.UserId = userDetailsType.ID
	if userStruct.UserId > 0 && userStruct.DocPicName != "" && userStruct.DocName != "" {
		log.Println("userStruct: ", userStruct)
		tx := utility.Db.MustBegin()
		inserted, err := models.KycDetails{}.Postkyc(userStruct, tx)
		if inserted && err == nil {
			err = tx.Commit()
			if err != nil {
				log.Println(err)
				tx.Rollback()
				response.Status = "400"
				response.Message = "Unable to update kyc at the moment! Please try again."
				utility.RenderJsonResponse(w, r, response, 400)
				return
			}
			response.Status = "200"
			response.Message = "Document upload successfully"
			utility.RenderJsonResponse(w, r, response, 200)
			return
		} else {
			tx.Rollback()
			response.Status = "400"
			response.Message = "Unable to update kyc at the moment! Please try again."
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
	} else {
		response.Status = "400"
		response.Message = "Please provide all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
}
func GetKycDetails(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	isok, userDetailsType := utility.CheckTokenPayloadAndReturnUser(r)
	if !isok {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	kyc_data, err := models.KycDetails{}.Getkyc(userDetailsType.ID)
	if err != nil {
		response.Status = "400"
		response.Message = "Unable to get kyc at the moment! Please try again."
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	response.Status = "200"
	response.Message = "successfully getting the record."
	response.Payload = kyc_data
	utility.RenderJsonResponse(w, r, response, 200)

}
func UploadHandler(w http.ResponseWriter, r *http.Request) bool {
	var savePath string
	response := utility.AjaxResponce{Status: "failure", Payload: ""}
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		response.Message = "Method not allowed"
		utility.RenderJsonResponse(w, r, response, 400)
		return false
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // Max 10 MB file size
	if err != nil {
		log.Println(err)
		response.Message = "Failed to parse form"
		utility.RenderJsonResponse(w, r, response, 400)
		return false
	}
	//get the modulename for the differentiations of the modules
	modulename := r.FormValue("modulename")
	// Get the file from the form data
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
		response.Message = "Failed to retrieve image from form data"
		utility.RenderJsonResponse(w, r, response, 400)
		return false
	}
	defer file.Close()
	//to save the screenshots with randomstring name and in upload folder.
	if modulename == "marketpurchase" {
		filename, err := utility.RandomNameForImage(handler)
		if err != nil {
			log.Println(err)
			response.Message = "Failed to generate image name."
			utility.RenderJsonResponse(w, r, response, 400)
			return false
		}
		// Create a new file on the server
		savePath = filepath.Join("uploads", filename)
		err = os.MkdirAll("uploads", os.ModePerm) // Create the "uploads" directory if it doesn't exist
	} else {
		if modulename == "item" {
			filename, err := utility.RandomNameForImage(handler)
			if err != nil {
				log.Println(err)
				response.Message = "Failed to generate image name."
				utility.RenderJsonResponse(w, r, response, 400)
				return false
			}
			handler.Filename = filename
		}
		//item image can be save in the assets folder.
		savePath = filepath.Join("assets/images/item", handler.Filename)
		err = os.MkdirAll("assets/images/item", os.ModePerm)
	}
	if err != nil {
		log.Println(err)
		response.Message = "Failed to create uploads directory"
		utility.RenderJsonResponse(w, r, response, 400)
		return false
	}
	newFile, err := os.Create(savePath)
	if err != nil {
		log.Println(err)
		response.Message = "Failed to create file on server"
		utility.RenderJsonResponse(w, r, response, 400)
		return false
	}
	defer newFile.Close()

	// Copy the uploaded file to the new file on the server
	_, err = io.Copy(newFile, file)
	if err != nil {
		log.Println(err)
		response.Message = "Failed to save file on server"
		utility.RenderJsonResponse(w, r, response, 400)
		return false
	}
	// Construct the URL for the saved file
	// baseURL := utility.GetBaseURL(r)
	// fileURL := utility.ConstructFileURL(baseURL, savePath)
	response.Status = "200"
	response.Message = "successfully getting the record."
	response.Payload = savePath
	// utility.RenderJsonResponse(w, r, response, 200)
	jsonresponce, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}
	trimmedResponse := bytes.TrimSpace(jsonresponce)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	w.Write(trimmedResponse)
	return true
	// jsonResponse, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
	// 	return
	// }

	// // Set the appropriate Content-Type header
	// w.Header().Set("Content-Type", "application/json")

	// // Send the JSON response
	// // w.Write(jsonResponse)
	// w.Write([]byte(jsonResponse))
}
