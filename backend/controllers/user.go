package controllers

import (
	"log"
	"strconv"

	"net/http"
	"reakgo/models"
	"reakgo/utility"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

func GetUserData(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponse{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}
	var para models.UserCondition
	var usr models.User
	err := utility.ReturnUserDetails(r, &usr)
	if err != nil {
		log.Println(err)
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		Utility.RenderJsonResponse(w, r, response)
		return true
	}

	if usr.Id == 0 || usr.CompanyId == 0 {
		response.Status = "403"
		response.Message = "Unauthorized access, UserId or companyId doesn't match."
		Utility.RenderJsonResponse(w, r, response)
		return true
	}

	para.AccountType = usr.AccountType
	para.Id = int64(utility.StrToInt(r.URL.Query().Get("id")))               // take id for url
	para.CompanyId = int64(utility.StrToInt(r.URL.Query().Get("CompanyId"))) // take company_id for url
	birthday := strings.ToLower(r.URL.Query().Get("birthday"))               // take birthday for url
	if birthday == "today" {
		para.Dob = strconv.Itoa(int(time.Now().Unix()))
	}
	if para.Id != 0 && para.AccountType == "user" {
		para.Id = usr.Id
	}
	parameters := models.User{}.GetParaForFilterUser(para)
	result, err := models.User{}.GetUser(parameters)
	if err != nil {
		log.Println(err)
	} else if len(result) == 0 {
		response.Status = "200"
		response.Message = "No result were found for this search."
	} else {
		response.Status = "200"
		response.Message = "Returns all matching users."
		response.Payload = result // Set the user data in the response payload
	}
	utility.RenderJsonResponse(w, r, response)
	return false
}

func DeleteUserData(w http.ResponseWriter, r *http.Request) bool {
	response := utility.AjaxResponse{Status: "500", Message: "Server is currently unavailable.", Payload: []interface{}{}}

	userId := utility.StrToInt(r.URL.Query().Get("id"))

	if userId <= 0 {
		response.Status = "400"
		response.Message = "Bad request, Incorrect payload or call."
		utility.RenderJsonResponse(w, r, response)
		return true
	}

	//Add data in user table then show the error
	boolType, err := models.User{}.DeleteUser(userId)
	if !boolType || err != nil {
		log.Println(err)
		response.Status = "500"
		response.Message = "Internal server error, Any serious issues which cannot be recovered from."
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			// MySQL error code 1451 indicates a foreign key constraint
			if driverErr.Number == 1451 {
				response.Message = utility.GetSqlErrorString(err)
			}
		}
		utility.RenderJsonResponse(w, r, response)
		return true
	}
	response.Status = "200"
	response.Message = "User deleted successfully."
	utility.RenderJsonResponse(w, r, response)
	return false
}
