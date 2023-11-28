package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

func IsValidUserStruct(userStruct models.Users) bool {
	if userStruct.Name != "" && userStruct.Email != "" && userStruct.PasswordHash != "" && userStruct.DOB != "" && userStruct.AccountType != "" && userStruct.TwoFactorKey != "" && userStruct.TwoFactorRecoveryCode != "" && userStruct.Status != "" {
		return true
	}
	return false
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (new decoder)
	var userStruct models.Users
	err := Helper.StrictParseDataFromJson(r, &userStruct)
	log.Println("userStruct: ", userStruct)
	if err != nil {
		Helper.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	// date format check
	if !Helper.CheckDateFormat(userStruct.DOB) {
		// Utility.Logger(err)
		response.Status = "403"
		response.Message = "Date is not in format `yyyy-mm-dd`"
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	if !Helper.CheckEmailFormat(userStruct.Email) {
		// Utility.Logger(err)
		response.Status = "403"
		response.Message = "Please enter valid email address"
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	ok, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if !ok {
		Helper.Logger(err)
		response.Status = "403"
		response.Message = "You cannot register the company because you are not an owner."
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	userType := userDetails.AccountType
	if userType == "" {
		userType = "guest"
	}

	if userType == "guest" && userStruct.AccountType != "owner" {
		// Utility.Logger(err)
		response.Status = "403"
		response.Message = "You cannot register without owners invite"
		Helper.RenderJsonResponse(w, r, response, 403)
		return

	} else if userType == "user" {
		// Utility.Logger(err)
		response.Status = "403"
		response.Message = "You cannot register without owners invite"
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	} else {
		// tokenPayload check
		isok, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
		if isok {
			if userStruct.AccountType == "owner" && userDetails.AccountType == "owner" {
				response.Status = "403"
				response.Message = "You are already registered as an owner for this company! Kindly use your credentials and login"
				Helper.RenderJsonResponse(w, r, response, 403)
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
		Helper.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Either required fields are empty or contain invalid data type"
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	// hashing password(plain text) #1
	userStruct.PasswordHash, err = Helper.SaltPlainPassWord(userStruct.PasswordHash)
	if err != nil {
		Helper.Logger(err)
		response.Status = "400"
		response.Message = "Unable to create a strong encryption for you password at the moment! Please try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	// only for now random twofactorkey and recovery code (to be deleted when integrated with ORM)
	userStruct.TwoFactorRecoveryCode, err = Helper.GenerateRandomString(16)
	if err != nil {
		log.Println("error in generating random string for TwoFactorRecoveryCode")
	}
	userStruct.TwoFactorKey, err = Helper.GenerateRandomString(16)
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
		isok, errString := Helper.CheckSqlError(err, "") // dummy check
		if isok {
			log.Println(errString)

		}
		Helper.Logger(err)
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		response.Status = "400"
		response.Message = "Unable to create user at the moment! Please try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	response.Status = "200"
	response.Message = "User created successfully."
	Helper.RenderJsonResponse(w, r, response, 200)

}

func PostUser(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	var userStruct models.Users
	err := Helper.StrictParseDataFromJson(r, &userStruct)
	if err != nil {
		// utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		return
	}
	// date format check
	if !Helper.CheckDateFormat(userStruct.DOB) {
		// Utility.Logger(err)
		response.Status = "400"
		response.Message = "Date is not in format `yyyy-mm-dd`"
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	if !Helper.CheckEmailFormat(userStruct.Email) {
		// Utility.Logger(err)
		response.Status = "400"
		response.Message = "Please enter valid email address"
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	var userDetails models.Users
	// integrate token check and return if mismatch found with status 400
	isok, userDetailsType := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isok {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	flag := Helper.CopyFieldsBetweenDiffStructType(userDetailsType, &userDetails)
	if !flag {
		log.Println("error during copy data at: CopyFieldsBetweenDiffStructType")
		response.Status = "400"
		response.Message = "Unable to process data at the moment! Please try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	if userStruct.Email == "" {
		response.Status = "400"
		response.Message = "Bad request! Cannot update data because of missing unique identifier"
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	// added a company id check too
	if userDetails.ID != userStruct.ID && userDetails.AccountType == "user" || userDetails.CompanyID != userStruct.CompanyID && userDetails.AccountType != "super-admin" {
		response.Status = "403"
		response.Message = "Unauthorized access! You are not allowed to make this request"
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	// this step is to get user incase a user with higher access rights tries to update other user
	if userDetails.AccountType == "owner" || userDetails.AccountType == "super-admin" && userStruct.AccountType != "super-admin" && userStruct.Email != userDetails.Email {
		// we need to get user details now
		// row, err := models.Users{}.GetUserById(userStruct.ID)
		//switch to email id not with the primary id
		row, err := models.Authentication{}.GetUserByEmail(userStruct.Email)
		if err != nil {
			log.Println("error: ", err)
			Helper.Logger(err)
			response.Status = "400"
			response.Message = "Unable to get user-record at the moment! Please try again."
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		}
		userDetails = CopyStructValues(row)
		//else assigning userdetails the row we just bought
		// userDetails = row
	}
	// check for passwordHash if empty
	if userStruct.PasswordHash == "" {
		userStruct.PasswordHash = userDetails.PasswordHash
	} else {
		// run salting rounds on userStruct.PasswordHash
		// hashing password(plain text) #1
		userStruct.PasswordHash, err = Helper.SaltPlainPassWord(userStruct.PasswordHash)
		if err != nil {
			Helper.Logger(err)
			response.Status = "400"
			response.Message = "Unable to create a strong encryption for your password at the moment! Please try again."
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		}
	}
	// check for twoFactory key/recoverCode
	if userStruct.TwoFactorKey == "" && userStruct.TwoFactorRecoveryCode == "" {
		// only for now random twofactorkey and recovery code (to be deleted when integrated with ORM)
		userStruct.TwoFactorRecoveryCode, err = Helper.GenerateRandomString(16)
		if err != nil {
			log.Println("error in generating random string for TwoFactorRecoveryCode")
		}
		userStruct.TwoFactorKey, err = Helper.GenerateRandomString(16)
		if err != nil {
			log.Println("error in generating random string for TwoFactorKey")
		}
	}
	// flipping values from userDetails(from token) to userStruct(from req) which are empty in userStruct
	flag = Helper.FillEmptyFieldsForPostUpdate(userDetails, &userStruct)
	if !flag {
		Helper.Logger(err)
		log.Println("error during flipping data at: FillEmptyFieldsForPostUser")
		response.Status = "400"
		response.Message = "Unable to process data at the moment! Please try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	// fill it with updated data
	if IsValidUserStruct(userStruct) {
		// call the ORM update function to update the user details
		tx := utility.Db.MustBegin()
		updateRow, err := models.Users{}.PostUser(userStruct)
		if err != nil {
			response.Status = "400"
			log.Println("under err")
			response.Message = Helper.GetSqlErrorString(err)
			log.Println("after sql err")
			if !updateRow {
				response.Message = "Record couldn't be updated"
			}
			tx.Rollback()
			response.Payload = []interface{}{}
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		}
		response.Status = "200"
		response.Message = "Record successfully updated"
		response.Payload = []interface{}{userStruct}

		userData, err := models.Authentication{}.GetUserByEmail(userStruct.Email)
		if err != nil {
			tx.Rollback()
			response.Status = "400"
			response.Message = "Unable to hydrate cache! Please try again."
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		} else {
			//update the IsWizardComplete from personal to company
			if userData.IsWizardComplete == "personal" && userData.PasswordHash != "" {
				userData.IsWizardComplete = "company"
				boolType, err := models.Users{}.UpdateWizardStatus(userData, tx)
				if err != nil {
					tx.Rollback()
					response.Status = "400"
					response.Message = "cant update the wizard status at that moment."
					Helper.RenderJsonResponse(w, r, response, 400)
					return
				}
				if boolType {
					err = tx.Commit()
					if err != nil {
						log.Println(err)
						tx.Rollback()
						response.Status = "400"
						response.Message = "Unable to update details at the moment! Please try again."
						Helper.RenderJsonResponse(w, r, response, 400)
						return
					}
				} else {
					tx.Rollback()
					response.Status = "400"
					response.Message = "Unable to update details at the moment! Please try again."
					Helper.RenderJsonResponse(w, r, response, 400)
					return
				}
			} else { //if the wizard is already completed or more progress is above personal we didnt update that and move forward
				err = tx.Commit()
				if err != nil {
					log.Println(err)
					tx.Rollback()
					response.Status = "400"
					response.Message = "Unable to update user at the moment! Please try again."
					Helper.RenderJsonResponse(w, r, response, 400)
					return
				}

			}
		}
		jsonData, _ := json.Marshal(userData)
		// rehydrating the cache after the a successful update
		utility.Cache.Set(userData.Token, jsonData)
		Helper.RenderJsonResponse(w, r, response, 200)
		return
	}

}

func GetUserData(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	var para models.UserCondition

	isOk, usr := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isOk {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	if usr.ID == 0 || usr.CompanyID == 0 {
		response.Status = "403"
		response.Message = "Unauthorized access, UserId or companyId doesn't match."
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}

	para.AccountType = usr.AccountType
	para.ID = int64(Helper.StrToInt(r.URL.Query().Get("id")))               // take id for url
	para.CompanyID = int64(Helper.StrToInt(r.URL.Query().Get("CompanyId"))) // take company_id for url
	birthday := strings.ToLower(r.URL.Query().Get("birthday"))              // take birthday for url
	if birthday == "today" {
		para.DOB = strconv.Itoa(int(time.Now().Unix()))
	}
	if para.ID != 0 && para.AccountType == "user" {
		para.ID = usr.ID
	}
	parameters := models.Users{}.GetParaForFilterUser(para)
	result, err := models.Users{}.GetUser(parameters)
	if err != nil {
		log.Println(err)
	} else if len(result) == 0 {
		response.Status = "400"
		response.Message = "No result were found for this search. Either record is not present or you are not authorized to access this users data"
		Helper.RenderJsonResponse(w, r, response, 400)
		return false
	} else {
		response.Status = "200"
		response.Message = "Returns all matching users."
		response.Payload = result // Set the user data in the response payload
	}
	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}

func DeleteUserData(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}

	userId := Helper.StrToInt(r.URL.Query().Get("id"))
	if userId <= 0 {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Helper.RenderJsonResponse(w, r, response, 400)
		return true
	}

	isok, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isok || userDetails.AccountType == "user" {
		response.Status = "403"
		response.Message = "You are not authorized to make this request."
		Helper.RenderJsonResponse(w, r, response, 403)
		return true
	}
	if userDetails.ID != int64(userId) {
		// get user by userID
		userData, err := models.Users{}.GetUserById(int64(userId))
		if err != nil {
			response.Status = "400"
			response.Message = "User doesn't exists."
			Helper.RenderJsonResponse(w, r, response, 400)
			return true
		}
		// check if owner and the user he's trying to delete belongs to the same company
		if userData.CompanyID != userDetails.CompanyID && userDetails.AccountType == "owner" {
			response.Status = "403"
			response.Message = "You are not authorized to make this request."
			Helper.RenderJsonResponse(w, r, response, 403)
			return true
		}
	}
	//Add data in user table then show the error
	boolType, err := models.Users{}.DeleteUser(int64(userId))
	if !boolType || err != nil {
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			// MySQL error code 1451 indicates a foreign key constraint
			if driverErr.Number == 1451 {
				response.Message = Helper.GetSqlErrorString(err)
			}
		}
		Helper.RenderJsonResponse(w, r, response, 500)
		return true
	}
	response.Status = "200"
	response.Message = "User deleted successfully."
	Helper.RenderJsonResponse(w, r, response, 200)
	return false
}
func CopyStructValues(st1 models.Authentication) models.Users {
	st2 := models.Users{
		ID:                    int64(st1.ID),
		Name:                  st1.Name,
		Email:                 st1.Email,
		PasswordHash:          st1.PasswordHash,
		TwoFactorKey:          st1.TwoFactorKey,
		TwoFactorRecoveryCode: st1.TwoFactorRecoveryCode,
		DOB:                   st1.DOB,
		AccountType:           st1.AccountType,
		CompanyID:             int64(st1.CompanyID),
		Status:                st1.Status,
		// Add other fields as needed
	}
	return st2
}
func PostUpdateWizardStatus(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	// wizardStatus := r.URL.Query().Get("wizardstatus")
	var user models.Authentication
	err := Helper.StrictParseDataFromJson(r, &user)
	if err != nil {
		Helper.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	if user.IsWizardComplete == "" {
		response.Status = "400"
		response.Message = "Please provide the wizard updated status and try again."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	isok, userDetails := Helper.CheckTokenPayloadAndReturnUser(r)
	if !isok {
		response.Status = "403"
		response.Message = "You are not authorized to make this request."
		Helper.RenderJsonResponse(w, r, response, 403)
		return
	}
	user.ID = userDetails.ID
	tx := utility.Db.MustBegin()
	boolType, err := models.Users{}.UpdateWizardStatus(user, tx)
	if err != nil {
		tx.Rollback()
		response.Status = "400"
		response.Message = "cant update the wizard status at the moment."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	if boolType {
		err = tx.Commit()
		if err != nil {
			log.Println(err)
			tx.Rollback()
			response.Status = "400"
			response.Message = "Unable to update wizard status at the moment! Please try again."
			Helper.RenderJsonResponse(w, r, response, 400)
			return
		}
		response.Status = "200"
		response.Message = "Document upload successfully"
		Helper.RenderJsonResponse(w, r, response, 200)
		return
	} else {
		tx.Rollback()
		response.Status = "400"
		response.Message = "cant update the wizard status at that moment."
		Helper.RenderJsonResponse(w, r, response, 400)
		return
	}
	response.Status = "400"
	response.Message = "cant update the wizard status at that moment."
	Helper.RenderJsonResponse(w, r, response, 400)
}
