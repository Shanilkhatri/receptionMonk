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
		utility.Logger(err)
		log.Println("Unable to decode json")
		response.Status = "400"
		response.Message = "Please check all fields correctly and try again."
		utility.RenderTemplate(w, r, "", response)
		return
	}
	// tokenPayload check
	isok, userDetails := utility.CheckTokenPayloadAndReturnUser(r)
	// only owner can enter call logs
	if isok && userDetails.AccountType == "owner" {
		if callLogsStruct.CallDuration != "" && callLogsStruct.CallExtension != "" && callLogsStruct.CallFrom != "" && callLogsStruct.CallPlacedAt != "" && callLogsStruct.CallTo != "" {
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
			}
			response.Status = "200"
			response.Message = "CallLog created successfully."
		}
	}
	utility.RenderTemplate(w, r, "", response)
}
