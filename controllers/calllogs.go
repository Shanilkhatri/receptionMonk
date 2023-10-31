package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
	"reakgo/utility"
)

func PutCallLogs(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	//decode json (later with ORM)
	var callLogsStruct models.CallLogs
	err := utility.StrictParseDataFromJson(r, &callLogsStruct)
	if err != nil {
		log.Println("Unable to decode json: ", err)
		// utility.Logger(err)
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderJsonResponse(w, r, response, 400)
		return
	}
	// tokenPayload check
	isok, userDetails := Utility.CheckTokenPayloadAndReturnUser(r)
	// only owner can enter call logs
	if !isok || userDetails.AccountType != "owner" {
		response.Status = "403"
		response.Message = "You are not authorized to make this request"
		utility.RenderJsonResponse(w, r, response, 403)
		return
	}
	// if companyId doesn't come as json we can set it here:
	// callLogsStruct.CompanyId = userDetails.CompanyID
	if callLogsStruct.CallDuration != "" && callLogsStruct.CallExtension != "" && callLogsStruct.CallFrom != "" && callLogsStruct.CallPlacedAt != "" && callLogsStruct.CallTo != "" && callLogsStruct.CompanyId != 0 {
		tx := utility.Db.MustBegin()
		isOk := models.CallLogs{}.PutCallLogs(callLogsStruct, tx)
		if !isOk {
			response.Status = "400"
			response.Message = "Unable to create CallLog at the moment! Please try again."
			isok, errString := utility.CheckSqlError(err, "") // dummy check
			if isok {
				log.Println(errString)
			}
			utility.Logger(err)
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			response.Status = "400"
			response.Message = "Unabke to create callLog at the moment. Please try again"
			utility.RenderJsonResponse(w, r, response, 400)
			return
		}
		response.Status = "200"
		response.Message = "CallLog created successfully."
		utility.RenderJsonResponse(w, r, response, 200)
		return
	}
}

func GetCallLogsDetails(w http.ResponseWriter, r *http.Request) {
	response := utility.AjaxResponce{Status: "500", Message: "Internal server error, Any serious issues which cannot be recovered from.", Payload: []interface{}{}}
	isOk, userDetails := Utility.CheckTokenPayloadAndReturnUser(r)
	log.Println("userDetails: ", userDetails)
	if isOk {
		queryParams := r.URL.Query()

		paramMap := map[string]int64{
			"id":        0,
			"companyId": 0,
		}
		for paramName := range paramMap {
			paramValue := queryParams.Get(paramName)
			if paramValue != "" {
				paramParsed, err := utility.StrToInt64(paramValue)
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
			Utility.RenderJsonResponse(w, r, response, 403)
			return
		}
		if paramMap["companyId"] != 0 && paramMap["companyId"] != userDetails.CompanyID {
			response.Status = "403"
			response.Message = "Unauthorized access! You are not authorized to make this request."
			Utility.RenderJsonResponse(w, r, response, 403)
			return
		}
		param := models.CallLogsCondition{
			CallLogs: models.CallLogs{
				Id:        paramMap["id"],
				CompanyId: paramMap["companyId"],
			},
		}

		parameters := models.CallLogs{}.GetParamForFilterCalllogs(param)
		result, err := models.CallLogs{}.GetCallLogs(parameters)
		if err != nil {
			log.Println(err)
			response.Status = "500"
			response.Message = "Internal server error, Any serious issues which cannot be recovered from."
			Utility.RenderJsonResponse(w, r, response, 500)
			return
		}

		if len(result) == 0 {
			response.Status = "200"
			response.Message = "No result were found for this search."
			Utility.RenderJsonResponse(w, r, response, 200)
			return
		} else {
			response.Status = "200"
			response.Message = "Returns all matching  calllogs."
			response.Payload = result // Set the calllogs data in the response payload
			Utility.RenderJsonResponse(w, r, response, 200)
			return
		}

	}
}
