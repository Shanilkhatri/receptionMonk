package controllers

import (
	"log"
	"net/http"
	"os"
	"reakgo/models"
	"reakgo/utility"
)

func IsValidUserStruct(userStruct models.Users) bool {
	if userStruct.Name != "" && userStruct.Email != "" && userStruct.PasswordHash != "" && userStruct.DOB != "" && userStruct.CompanyID != 0 && userStruct.AccountType != "" && userStruct.TwoFactorKey != "" && userStruct.TwoFactorRecoveryCode != "" && userStruct.Status != "" && userStruct.ID == 0 {
		return true
	}
	return false
}
func PutUser(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var userStruct models.Users
	err := utility.StrictParseDataFromJson(r, &userStruct)
	if err != nil {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderTemplate(w, r, "", response)
		return
	}
	if CheckACL(w, r, []string{"user", "owner", "super-admin"}) {
		// tokenPayload check
		isok, userDetails := utility.CheckTokenPayloadAndReturnUser(r)
		if isok {
			if userStruct.AccountType == "owner" {
				response.Status = "403"
				response.Message = "You are already registered as an owner for this company! Kindly use your credentials and login"
				utility.RenderTemplate(w, r, "", response)
				return
			}
			// register as a user
			// takeout the company id from owners token
			log.Println("owner Details: ", userDetails)
			companyId := userDetails.CompanyID // which we'll get from the GET Op

			// setting company id
			userStruct.CompanyID = companyId

		}
	}
	// make an entry into cpmay table when AccountType == "owner"
	if userStruct.AccountType == "owner" {
		// call PUT company here
		// if successfull then GET comapny id and store it into
		// userStruct.CompanyId
		// in case of error response.Status = "403" response.Message = "Unable to Process request at the moment"

		// dummy company id
		userStruct.CompanyID = 2
	}
	if IsValidUserStruct(userStruct) {
		// hashing password(plain text) #1
		userStruct.PasswordHash, err = utility.NewPasswordHash(userStruct.PasswordHash)
		if err != nil {
			// handle the error
			utility.Logger(err)
			response.Status = "400"
			response.Message = "Unable to create user at the moment! Please try again."

		}
		// mixing salt with hashed pass
		pswdConcatWithSalt := userStruct.PasswordHash + os.Getenv("CONS_SALT")

		// making hash of (salted+hashed) pass #2
		userStruct.PasswordHash, err = utility.NewPasswordHash(pswdConcatWithSalt)
		if err != nil {
			// handle the error
			utility.Logger(err)
			response.Status = "400"
			response.Message = "Unable to create user at the moment! Please try again."
		}
		// only for now random twofactorkey and recovery code (to be deleted when integrated with ORM)
		userStruct.TwoFactorRecoveryCode, err = utility.GenerateRandomString(16)
		if err != nil {
			log.Println("error in generating random string for TwoFactorRecoveryCode")
		}
		userStruct.TwoFactorKey, err = utility.GenerateRandomString(16)
		if err != nil {
			log.Println("error in generating random string for TwoFactorKey")
		}
		// ^^^only for now random twofactorkey and recovery code (to be deleted when integrated with ORM)^^^
		log.Println("userStruct: ", userStruct)
		tx := utility.Db.MustBegin()
		isok := models.Users{}.PutUser(userStruct, tx)
		if !isok {
			response.Status = "400"
			response.Message = "Unable to create user at the moment! Please try again."
			isok, errString := utility.CheckSqlError(err, "") // dummy check
			if isok {
				log.Println(errString)

			}
			tx.Rollback()
			utility.Logger(err)
			utility.RenderTemplate(w, r, "", response)
			return
		}
		response.Status = "200"
		response.Message = "User created successfully."
	}
	utility.RenderTemplate(w, r, "", response)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}

	var userStruct models.Users
	err := utility.StrictParseDataFromJson(r, &userStruct)
	if err != nil {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."

	}
	// integrate token check and return if mismatch found with status 400
	isok, userDetails := utility.CheckTokenPayloadAndReturnUser(r)
	if isok {
		if userStruct.ID == 0 {
			response.Status = "400"
			response.Message = "Bad request! Cannot update data because of missing unique identifier"
			utility.RenderTemplate(w, r, "", response)
			return
		}
		if userDetails.ID != userStruct.ID {
			response.Status = "403"
			response.Message = "Unauthorized access! You are not allowed to make this request"
			utility.RenderTemplate(w, r, "", response)
			return
		}

		// fill it with updated data
		if userStruct.Name != "" && userStruct.Email != "" && userStruct.PasswordHash != "" && userStruct.DOB != "" && userStruct.CompanyID != 0 && userStruct.AccountType != "" && userStruct.TwoFactorKey != "" && userStruct.TwoFactorRecoveryCode != "" && userStruct.Status != "" && userStruct.ID != 0 {
			// call the ORM update function to update the user details
			log.Println("Begin post User...")
			//Backend data update and pass payload map
			updateRow, err := models.Users{}.PostUser(userStruct)
			if err != nil {
				response.Message = utility.GetSqlErrorString(err)
				if !updateRow {
					response.Message = "400"
					response.Message = "Record couldn't be updated"
					response.Payload = []interface{}{}
				}
			} else {
				response.Status = "200"
				response.Message = "Record successfully updated"
				response.Payload = []interface{}{userStruct}
			}
		}
	}
	utility.RenderTemplate(w, r, "", response)
}
