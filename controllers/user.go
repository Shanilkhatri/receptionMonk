package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
)

func IsValidUserStruct(userStruct models.Users) bool {
	if userStruct.Name != "" && userStruct.Email != "" && userStruct.PasswordHash != "" && userStruct.DOB != "" && userStruct.CompanyID != 0 && userStruct.AccountType != "" && userStruct.TwoFactorKey != "" && userStruct.TwoFactorRecoveryCode != "" && userStruct.Status != "" {
		return true
	}
	return false
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var userStruct models.Users
	err := utility.StrictParseDataFromJson(r, &userStruct)
	log.Println("userStruct: ", userStruct)
	if err != nil {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// date format check
	if !utility.CheckDateFormat(userStruct.DOB) {
		// Utility.Logger(err)
		response.Status = "403"
		response.Message = "Date is not in format `yyyy-mm-dd`"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	if !utility.CheckEmailFormat(userStruct.Email) {
		// Utility.Logger(err)
		response.Status = "403"
		response.Message = "Please enter valid email address"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	userType := Utility.SessionGet(r, "type")
	if userType == nil {
		userType = "guest"
	}

	if userType == "guest" && userStruct.AccountType != "owner" {
		// Utility.Logger(err)
		response.Status = "403"
		response.Message = "You cannot register without owners invite"
		utility.RenderJsonResponse(w, r, response, 403)
		return

	} else if userType == "user" {
		// Utility.Logger(err)
		response.Status = "403"
		response.Message = "You cannot register without owners invite"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	} else {
		// tokenPayload check
		isok, userDetails := Utility.CheckTokenPayloadAndReturnUser(r)
		if isok {
			if userStruct.AccountType == "owner" && userDetails.AccountType == "owner" {
				response.Status = "403"
				response.Message = "You are already registered as an owner for this company! Kindly use your credentials and login"
				utility.RenderJsonResponse(w, r, response, 403)
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
	if !IsValidUserStruct(userStruct) {
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Either required fields are empty or contain invalid data type"
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// hashing password(plain text) #1
	userStruct.PasswordHash, err = utility.SaltPlainPassWord(userStruct.PasswordHash)
	if err != nil {
		utility.Logger(err)
		response.Status = "400"
		response.Message = "Unable to create a strong encryption for you password at the moment! Please try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// only for now random twofactorkey and recovery code (to be deleted when integrated with ORM)
	userStruct.TwoFactorRecoveryCode, err = Utility.GenerateRandomString(16)
	if err != nil {
		log.Println("error in generating random string for TwoFactorRecoveryCode")
	}
	userStruct.TwoFactorKey, err = Utility.GenerateRandomString(16)
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
		utility.Logger(err)
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		response.Status = "400"
		response.Message = "Unable to create user at the moment! Please try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	response.Status = "200"
	response.Message = "User created successfully."
	utility.RenderJsonResponse(w, r, response, 200)

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
	var userDetails models.Users
	// integrate token check and return if mismatch found with status 400
	isok, userDetailsType := Utility.CheckTokenPayloadAndReturnUser(r)
	flag := utility.CopyFieldsBetweenDiffStructType(userDetailsType, &userDetails)
	if !flag {
		log.Println("error during copy data at: CopyFieldsBetweenDiffStructType")
		response.Status = "400"
		response.Message = "Unable to process data at the moment! Please try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	if !isok {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	if userStruct.ID == 0 {
		response.Status = "400"
		response.Message = "Bad request! Cannot update data because of missing unique identifier"
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// added a company id check too
	if userDetails.ID != userStruct.ID && userDetails.AccountType == "user" || userDetails.CompanyID != userStruct.CompanyID && userDetails.AccountType != "super-admin" {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	// check for passwordHash if empty
	if userStruct.PasswordHash == "" {
		userStruct.PasswordHash = userDetails.PasswordHash
	} else {
		// run salting rounds on userStruct.PasswordHash
		// hashing password(plain text) #1
		userStruct.PasswordHash, err = utility.SaltPlainPassWord(userStruct.PasswordHash)
		if err != nil {
			utility.Logger(err)
			response.Status = "400"
			response.Message = "Unable to create a strong encryption for your password at the moment! Please try again."
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
	}
	// check for twoFactory key/recoverCode
	if userStruct.TwoFactorKey == "" && userStruct.TwoFactorRecoveryCode == "" {
		// only for now random twofactorkey and recovery code (to be deleted when integrated with ORM)
		userStruct.TwoFactorRecoveryCode, err = Utility.GenerateRandomString(16)
		if err != nil {
			log.Println("error in generating random string for TwoFactorRecoveryCode")
		}
		userStruct.TwoFactorKey, err = Utility.GenerateRandomString(16)
		if err != nil {
			log.Println("error in generating random string for TwoFactorKey")
		}
	}
	// flipping values from userDetails(from token) to userStruct(from req) which are empty in userStruct
	flag = utility.FillEmptyFieldsForPostUser(userDetails, &userStruct)
	if !flag {
		utility.Logger(err)
		log.Println("error during flipping data at: FillEmptyFieldsForPostUser")
		response.Status = "400"
		response.Message = "Unable to process data at the moment! Please try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// fill it with updated data
	if IsValidUserStruct(userStruct) {
		// call the ORM update function to update the user details
		log.Println("Begin post User...")
		tx := utility.Db.MustBegin()
		updateRow, err := models.Users{}.PostUser(userStruct)
		if err != nil {
			response.Status = "400"
			log.Println("under err")
			response.Message = utility.GetSqlErrorString(err)
			log.Println("after sql err")
			if !updateRow {
				response.Message = "Record couldn't be updated"
			}
			tx.Rollback()
			response.Payload = []interface{}{}
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Println(err)
			tx.Rollback()
			response.Status = "400"
			response.Message = "Unable to update user at the moment! Please try again."
			utility.RenderJsonResponse(w, r, response, 400)
			return

		}
		response.Status = "200"
		response.Message = "Record successfully updated"
		response.Payload = []interface{}{userStruct}
		utility.RenderJsonResponse(w, r, response, 200)
		return
	}

}
