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
	err := Helper.StrictParseDataFromJson(r, &walletStruct)
	if err != nil {
		Helper.Logger(err)
		response.Status = "400"
		response.Message = "Please fill all the fields correctly and try again"
		Helper.RenderJsonResponse(w, r, response, 400)
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
				Helper.RenderJsonResponse(w, r, response, 500)
			}
		}()
		_, err := models.Wallet{}.PutWallet(walletStruct, tx)
		if err != nil {
			sqlErr := Helper.GetSqlErrorString(err)
			response.Status = "400"
			response.Message = "Couldn't add wallet info at the moment! Please try again."
			if sqlErr != "" {
				response.Message = sqlErr
			}
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		}
		tx.Commit()
		response.Status = "200"
		response.Message = "Wallet information added successfully."
		Helper.RenderJsonResponse(w, r, response, 200)
		return
	}
	response.Status = "400"
	response.Message = "Please fill the required fields properly, leaving them vacant will result in non-submission."
	Helper.RenderJsonResponse(w, r, response, 400)
}

func PostWallet(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Please wait a moment, the server is currently unavailable.", Payload: []interface{}{}}

	var walletStruct models.Wallet
	// directly unmarshalling the json data into the struct
	err := Helper.StrictParseDataFromJson(r, &walletStruct)
	if err != nil {
		Helper.Logger(err)
		response.Status = "400"
		response.Message = "Please fill all the fields correctly and try again"
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}

	// getting userDetails from the token
	isOk, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	// as directed in the documentation, only accountType "owner" can access this route
	if !isOk || userDetails.AccountType != "owner" {
		response.Status = "403"
		response.Message = "You are not authorized to make this request."
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	// if the owner is trying to update wallet belonging to a diff company
	// -> we'll get the wallet by id
	walletData, _ := models.Wallet{}.GetWalletById(walletStruct.Id)
	// performing check
	if walletData.CompanyId != userDetails.CompanyID {
		response.Status = "403"
		response.Message = "You are not authorized to make this request."
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	// updating epoch to mark updated time
	walletStruct.Epoch = time.Now().Unix()
	// copying all the unchanged fields from src to dest
	flag := Helper.FillEmptyFieldsForPostUpdate(walletData, &walletStruct)
	if !flag {
		log.Println("couldn't flip fields for wallet")
		response.Status = "400"
		response.Message = "Unable to update data at the moment! Please try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	// now calling the model function to finally update the wallet data at DB
	_, err = models.Wallet{}.PostWallet(walletStruct)
	if err != nil {
		Helper.Logger(err)
		sqlErr := Helper.GetSqlErrorString(err)
		response.Status = "400"
		response.Message = "Couldn't update wallet info at the moment! Please try again."
		if sqlErr != "" {
			response.Message = sqlErr
		}
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	response.Status = "200"
	response.Message = "Wallet information updated successfully."
	Helper.RenderJsonResponse(w, r, response, 200)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	isOk, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)

	if isOk {
		queryParams := r.URL.Query()

		paramMap := map[string]int64{
			"id":        0,
			"companyId": 0,
		}
		for paramName := range paramMap {
			paramValue := queryParams.Get(paramName)
			if paramValue != "" {
				paramParsed, err := Helper.StrToInt64(paramValue)
				if err != nil {
					log.Println(err)
				} else {
					paramMap[paramName] = paramParsed
				}
			}
		}

		// Check if the user is authorized to access call logs
		if userDetails.AccountType != "owner" {
			response.Status = "403"
			response.Message = "Unauthorized access! You are not authorized to make this request."
			Helper.RenderJsonResponse(w, r, response, 403)
			return
		}
		if paramMap["companyId"] != 0 && paramMap["companyId"] != userDetails.CompanyID {
			response.Status = "403"
			response.Message = "Unauthorized access! You are not authorized to make this request."
			Helper.RenderJsonResponse(w, r, response, 403)
			return
		}
		param := models.WalletCondition{
			Wallet: models.Wallet{
				Id:        paramMap["id"],
				CompanyId: paramMap["companyId"],
			},
		}

		parameters := models.Wallet{}.GetParamForFilterWallet(param)
		result, err := models.Wallet{}.GetWallet(parameters)
		if err != nil {
			log.Println(err)
			response.Status = "500"
			response.Message = "Internal server error, Any serious issues which cannot be recovered from."
			Helper.RenderJsonResponse(w, r, response, 500)
			return
		}

		if len(result) == 0 {
			response.Status = "200"
			response.Message = "No result were found for this search."
			Helper.RenderJsonResponse(w, r, response, 200)
			return
		} else {
			response.Status = "200"
			response.Message = "Fetched wallet transactions successfully."
			response.Payload = result // Set the calllogs data in the response payload
			Helper.RenderJsonResponse(w, r, response, 200)
			return
		}

	}
	response.Status = "403"
	response.Message = "You are not authorized to make this request."
	Helper.RenderJsonResponse(w, r, response, 403)

}
func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}

	walletId := Helper.StrToInt(r.URL.Query().Get("id"))
	if walletId <= 0 {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	isok, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isok || userDetails.AccountType != "owner" && userDetails.AccountType != "super-admin" {
		response.Status = "403"
		response.Message = "You are not authorized to make this request."
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	if userDetails.AccountType == "owner" {
		// get the ticket he wants to delete and check if the userId matches
		walletData, err := models.Wallet{}.GetWalletById(int64(walletId))
		if err != nil {
			response.Status = "400"
			response.Message = Helper.GetSqlErrorString(err)
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		}
		if walletData.CompanyId != userDetails.CompanyID {
			response.Status = "403"
			response.Message = "You are not authorized to make this request."
			Helper.RenderJsonResponse(w, r, response, 403)
			return
		}
	}
	_, err := models.Wallet{}.DeleteWallet(int64(walletId))
	if err != nil {
		response.Status = "400"
		response.Message = Helper.GetSqlErrorString(err)
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	response.Status = "200"
	response.Message = "Wallet deleted successfully."
	Helper.RenderJsonResponse(w, r, response, 200)
}
