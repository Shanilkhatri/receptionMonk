package models

import (
	"log"
	"reakgo/utility"

	"github.com/jmoiron/sqlx"
)

type Response struct {
	Id           int64  `db:"id" json:"id"`
	Response     string `db:"response" json:"response"`
	TicketId     int64  `db:"ticketId" json:"ticketId"`
	ResponseTime int64  `db:"responseTime" json:"responseTime"`
	Type         string `db:"type" json:"type"`
	RespondeeId  int64  `db:"respondeeId" json:"respondeeId"`
}
type ResponseCondition struct {
	WhereCondition string
	Response
}

func (Response) PutResponse(responseStruct Response, tx *sqlx.Tx) bool {
	// executing query
	_, err := tx.NamedExec("INSERT INTO `responses` (response,ticketId,responseTime,type,respondeeId) VALUES ( :Response,:TicketId,:ResponseTime,:Type,:RespondeeId)", map[string]interface{}{"Response": responseStruct.Response, "TicketId": responseStruct.TicketId, "ResponseTime": responseStruct.ResponseTime, "Type": responseStruct.Type, "RespondeeId": responseStruct.RespondeeId})
	// Check error
	if err != nil {
		log.Println(err)
		//logger remove for duplicate entry that means duplicate error message not send at email.
		istrue, _ := utility.CheckSqlError(err, "Duplicate entry")
		if !istrue {
			utility.Logger(err)
		}
		return false
	} else {
		return true
	}
}

func (Response) GetResponses(ticketId int64) ([]Response, error) {
	var responseData []Response
	query := "SELECT * from responses Where 1=1 AND ticketId=:TicketId ORDER BY responseTime DESC;"
	condition := map[string]interface{}{
		"TicketId": ticketId,
	}
	rows, err := utility.Db.NamedQuery(query, condition)
	if err != nil {
		log.Println(err)
		return responseData, err
	}
	defer rows.Close()
	for rows.Next() {
		var singleRow Response
		err := rows.Scan(&singleRow.Id, &singleRow.Response, &singleRow.TicketId, &singleRow.ResponseTime, &singleRow.Type, &singleRow.RespondeeId)
		if err != nil {
			log.Println(err)
			return responseData, err
		}
		responseData = append(responseData, singleRow)
	}

	return responseData, nil
}
