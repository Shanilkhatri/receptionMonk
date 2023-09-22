package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reakgo/models"
	"reakgo/utility"
)

func UserJsonDecoder(r *http.Request) (models.Users, error) {
	var userStruct models.Users
	err := json.NewDecoder(r.Body).Decode(&userStruct)
	return userStruct, err
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (later with ORM)
	userStruct, err := UserJsonDecoder(r)
	if err != nil {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."

	}
	if userStruct.Name != "" && userStruct.Email != "" && userStruct.PasswordHash != "" && userStruct.DOB != "" && userStruct.CompanyID != 0 && userStruct.AccountType != "" && userStruct.TwoFactorKey != "" && userStruct.TwoFactorRecoveryCode != "" && userStruct.Status != "" {
		// mixing salt with pass
		pswdConcatWithSalt := userStruct.PasswordHash + os.Getenv("CONS_SALT")
		// making hash of salted pass
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
		err = Db.users.PutUser("", userStruct)
		if err != nil {
			response.Status = "400"
			response.Message = "Unable to create user at the moment! Please try again."
			isok, errString := utility.CheckSqlError(err, "") // dummy check
			if isok {
				log.Println(errString)

			}
			utility.Logger(err)
		}
		response.Status = "200"
		response.Message = "User created successfully."
	}
	utility.RenderTemplate(w, r, "", response)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	// integrate token check and return if mismatch found with status 400
	//decode json (later with ORM)
	userStruct, err := UserJsonDecoder(r)
	if err != nil {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."

	}
	if userStruct.ID == 0 {
		response.Status = "400"
		response.Message = "Bad request! Cannot update data because of missing unique identifier"
		utility.RenderTemplate(w, r, "", response)
		return
	}
	// use find function(ORM) here to fetch all the records using primary key
	// empty fields in userStruct means they are not changed so update thier values with the values that find() returned

	// fill it with updated data
	if userStruct.Name != "" && userStruct.Email != "" && userStruct.PasswordHash != "" && userStruct.DOB != "" && userStruct.CompanyID != 0 && userStruct.AccountType != "" && userStruct.TwoFactorKey != "" && userStruct.TwoFactorRecoveryCode != "" && userStruct.Status != "" && userStruct.ID != 0 {
		// call the ORM update function to update the user details
	}
}
