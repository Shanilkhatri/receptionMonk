package models

import (
	"log"
	"reakgo/utility"

	"github.com/jmoiron/sqlx"
)

type CallLogs struct {
	Id            int64  `json:"id" db:"id"`
	CallFrom      string `json:"callFrom" db:"callFrom"`
	CallTo        string `json:"callTo" db:"callTo"`
	CallPlacedAt  string `json:"callPlacedAt" db:"callPlacedAt"`
	CallDuration  string `json:"callDuration" db:"callDuration"`
	CallExtension string `json:"callExtension" db:"callExtension"`
	CompanyId     int64  `json:"companyId" db:"companyId"`
}

type CallLogsCondition struct {
	CallLogs
	WhereCondition string
}

// func (callLogs CallLogs) PutCallLogs(tablename string, structure CallLogs) error {
// 	// sql INSERTION with ORM here
// 	return nil
// }

func (callLogs CallLogs) PutCallLogs(logs CallLogs, tx *sqlx.Tx) bool {

	_, err := tx.NamedExec("INSERT INTO callLogs (callFrom,callTo,callPlacedAt,callDuration,callExtension,companyId) VALUES ( :CallFrom,:CallTo,:CallPlacedAt,:CallDuration,:CallExtension,:CompanyId)", map[string]interface{}{"CallFrom": logs.CallFrom, "CallTo": logs.CallTo, "CallPlacedAt": logs.CallPlacedAt, "CallDuration": logs.CallDuration, "CallExtension": logs.CallExtension, "CompanyId": logs.CompanyId})
	// Check error
	if err != nil {
		log.Println(err)
		//logger remove for duplicate entry that means duplicate error message not send at email.
		istrue, _ := Helper.CheckSqlError(err, " Duplicate entry")
		if !istrue {
			Helper.Logger(err)
		}
		return false
	} else {
		return true
	}

}

func (callLogs CallLogs) GetCallLogs(filter CallLogsCondition) ([]CallLogs, error) {
	var calllogsData []CallLogs
	query := "SELECT * FROM `calllogs` WHERE 1=1" + filter.WhereCondition
	condition := map[string]interface{}{
		"Id":        filter.Id,
		"CompanyId": filter.CompanyId,
	}
	rows, err := utility.Db.NamedQuery(query, condition)

	if err != nil {
		log.Println(err)
		return calllogsData, err
	}
	defer rows.Close()
	for rows.Next() {
		var singleRow CallLogs
		err := rows.Scan(&singleRow.Id, &singleRow.CallFrom, &singleRow.CallTo, &singleRow.CallPlacedAt, &singleRow.CallDuration, &singleRow.CallExtension, &singleRow.CompanyId)
		if err != nil {
			log.Println(err)
			return calllogsData, err
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
		calllogsData = append(calllogsData, singleRow)
	}
	return calllogsData, err
}

func (callLogs CallLogs) GetParamForFilterCalllogs(param CallLogsCondition) CallLogsCondition {
	if param.Id != 0 {
		param.WhereCondition += " AND id=:Id "
	}
	if param.CompanyId != 0 {
		param.WhereCondition += " AND companyId=:CompanyId "
	}

	return param
}
