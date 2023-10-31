package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
	"time"
)

func PutWallet(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Please wait a moment, the server is currently unavailable.", Payload: []interface{}{}}

	// there are no checks here as the documentation says it needs to bypass header auth
	var walletStruct models.Wallet
	// directly unmarshalling the json data into the struct
	err := utility.StrictParseDataFromJson(r, &walletStruct)
	if err != nil {
		utility.Logger(err)
		response.Status = "400"
		response.Message = "Please fill all the fields correctly and try again"
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// populating epoch with current time in unix timestamp
	walletStruct.Epoch = time.Now().Unix()
	// checking for : if empty struct fields
	if walletStruct.Charge != "" && walletStruct.CompanyId != 0 && walletStruct.Cost != 0 && walletStruct.Epoch != 0 && walletStruct.Reason != "" {
		tx := utility.Db.MustBegin()
		// for handling any panic
		defer func() {
			if recover := recover(); recover != nil {
				log.Println("panic occured: ", recover)
				tx.Rollback()
				response.Message = "An internal error occurred, please try again"
				utility.RenderJsonResponse(w, r, response, 500)
			}
		}()
		_, err := models.Wallet{}.PutWallet(walletStruct, tx)
		if err != nil {
			sqlErr := utility.GetSqlErrorString(err)
			response.Status = "400"
			response.Message = "Couldn't add wallet info at the moment! Please try again."
			if sqlErr != "" {
				response.Message = sqlErr
			}
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
		tx.Commit()
		response.Status = "200"
		response.Message = "Wallet information added successfully."
		utility.RenderJsonResponse(w, r, response, 200)
		return
	}
	response.Status = "400"
	response.Message = "Please fill the required fields properly, leaving them vacant will result in non-submission."
	utility.RenderJsonResponse(w, r, response, 400)
}

func PostWallet(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Please wait a moment, the server is currently unavailable.", Payload: []interface{}{}}

	var walletStruct models.Wallet
	// directly unmarshalling the json data into the struct
	err := utility.StrictParseDataFromJson(r, &walletStruct)
	if err != nil {
		utility.Logger(err)
		response.Status = "400"
		response.Message = "Please fill all the fields correctly and try again"
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}

	// getting userDetails from the token
	isOk, userDetails := utility.CheckTokenPayloadAndReturnUser(r)
	// as directed in the documentation, only accountType "owner" can access this route
	if !isOk || userDetails.AccountType != "owner" {
		response.Status = "403"
		response.Message = "You are not authorized to make this request."
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	// if the owner is trying to update wallet belonging to a diff company
	// -> we'll get the wallet by id
	walletData, _ := models.Wallet{}.GetWalletById(walletStruct.Id)
	// performing check
	if walletData.CompanyId != userDetails.CompanyID {
		response.Status = "403"
		response.Message = "You are not authorized to make this request."
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	// updating epoch to mark updated time
	walletStruct.Epoch = time.Now().Unix()
	// copying all the unchanged fields from src to dest
	flag := utility.FillEmptyFieldsForPostUpdate(walletData, &walletStruct)
	if !flag {
		log.Println("couldn't flip fields for wallet")
		response.Status = "400"
		response.Message = "Unable to update data at the moment! Please try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// now calling the model function to finally update the wallet data at DB
	_, err = models.Wallet{}.PostWallet(walletStruct)
	if err != nil {
		utility.Logger(err)
		sqlErr := utility.GetSqlErrorString(err)
		response.Status = "400"
		response.Message = "Couldn't update wallet info at the moment! Please try again."
		if sqlErr != "" {
			response.Message = sqlErr
		}
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	response.Status = "200"
	response.Message = "Wallet information updated successfully."
	utility.RenderJsonResponse(w, r, response, 200)
}
