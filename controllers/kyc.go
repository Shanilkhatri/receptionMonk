package controllers

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reakgo/models"
	"reakgo/utility"
	"strconv"
	"strings"
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

// curl -X POST -H "Authorization: Bearer $2a$10$e3hXRBp5LaepakdYrq2RFegovRN9ivfiVBqF49qC0m6hIcyRfB1Zm" -F "image=@/home/user/Pictures/Screenshots/Screenshot from 2023-10-16 11-24-16 .png" -F "modulename=kyc" http://localhost:4000/kycfileupload
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var savePath string
	response := utility.AjaxResponce{Status: "400", Payload: []interface{}{}}
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		response.Message = "Method not allowed"
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // Max 10 MB file size
	if err != nil {
		// utility.Logger(err)
		response.Message = "Failed to parse form"
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	//get the modulename for the differentiations of the modules
	modulename := r.FormValue("modulename")
	// Get the file from the form data
	file, handler, err := r.FormFile("image")
	if err != nil {
		// utility.Logger(err)
		response.Message = "Failed to retrieve image from form data"
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	defer file.Close()
	//to save the screenshots with randomstring name and in upload folder.
	if modulename == "kyc" {
		filename, err := RandomNameForImage(handler)
		if err != nil {
			// utility.Logger(err)
			response.Message = "Failed to generate image name."
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
		idString := strconv.FormatInt(userDetailsType.ID, 10)
		name, _ := utility.NewPasswordHash(idString + userDetailsType.Name)
		if name == "" {
			response.Message = "Failed to find the name for this image folder."
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
		folderName := strings.ReplaceAll(name, "/", "")
		err = os.MkdirAll("uploads", os.ModePerm) // Create the "uploads" directory if it doesn't exist
		randomFolderPath := filepath.Join("uploads", folderName)
		err = os.MkdirAll(randomFolderPath, os.ModePerm)
		if err != nil {
			// utility.Logger(err)
			response.Message = "Failed to save this image"
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
		// Create a new file on the server
		savePath = filepath.Join(randomFolderPath, filename)
	} else {
		if modulename == "item" {
			filename, err := RandomNameForImage(handler)
			if err != nil {
				// utility.Logger(err)
				response.Message = "Failed to generate image name."
				utility.RenderJsonResponse(w, r, response, 400)
				return
			}
			handler.Filename = filename
		}
		//item image can be save in the assets folder.
		savePath = filepath.Join("assets/images/item", handler.Filename)
		err = os.MkdirAll("assets/images/item", os.ModePerm)
	}
	if err != nil {
		// utility.Logger(err)
		response.Message = "Failed to create uploads directory"
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	newFile, err := os.Create(savePath)
	if err != nil {
		// utility.Logger(err)
		response.Message = "Failed to create file on server"
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	defer newFile.Close()

	// Copy the uploaded file to the new file on the server
	_, err = io.Copy(newFile, file)
	if err != nil {
		// utility.Logger(err)
		response.Message = "Failed to save file on server"
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// Construct the URL for the saved file
	baseURL := getBaseURL(r)
	fileURL := constructFileURL(baseURL, savePath)
	log.Println("Exact filepath:", fileURL)

	response.Status = "200"
	response.Message = "File uploaded successfully"
	response.Payload = savePath
	utility.RenderJsonResponse(w, r, response, 200)
}
func getBaseURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}

func constructFileURL(baseURL, filePath string) string {
	// Clean the file path to ensure it doesn't contain any leading slash
	filePath = strings.TrimLeft(filePath, "/")
	// Encode the file path to handle special characters
	encodedFilePath := url.PathEscape(filePath)
	// Concatenate the base URL and encoded file path to get the complete URL
	// baseURL + "/" + encodedFilePath
	replacedStr := strings.ReplaceAll(baseURL+"/"+encodedFilePath, "%2F", "/")
	return replacedStr
}
func RandomNameForImage(handler *multipart.FileHeader) (string, error) {
	var extension string
	//for getting the type of the image
	// lastDotIndex := strings.LastIndex(handler.Filename, ".")
	// if lastDotIndex != -1 {
	// 	extension = handler.Filename[lastDotIndex:]
	// }
	extension = utility.GetImageTypeExtension(handler.Filename, ".", true)
	randomString, err := utility.GenerateRandomString(30)
	if err != nil {
		utility.Logger(err)
		// response.Message = "Failed to generate image name."
		// utility.RenderJsonResponse(w, r, response, 400)
		return "", err
	}
	//filename with its extension.
	filename := randomString + extension
	return filename, err
}
