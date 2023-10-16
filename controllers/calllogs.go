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
