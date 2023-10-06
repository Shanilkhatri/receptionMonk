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
}

// func (callLogs CallLogs) PutCallLogs(tablename string, structure CallLogs) error {
// 	// sql INSERTION with ORM here
// 	return nil
// }

func (callLogs CallLogs) PutCallLogs(logs CallLogs, tx *sqlx.Tx) bool {

	_, err := tx.NamedExec("INSERT INTO callLogs (callFrom,callTo,callPlacedAt,callDuration,callExtension) VALUES ( :CallFrom,:CallTo,:CallPlacedAt,:CallDuration,:CallExtension)", map[string]interface{}{"CallFrom": logs.CallFrom, "CallTo": logs.CallTo, "CallPlacedAt": logs.CallPlacedAt, "CallDuration": logs.CallDuration, "CallExtension": logs.CallExtension})
	// Check error
	if err != nil {
		log.Println(err)
		//logger remove for duplicate entry that means duplicate error message not send at email.
		istrue, _ := utility.CheckSqlError(err, " Duplicate entry")
		if !istrue {
			utility.Logger(err)
		}
		return false
	} else {
		return true
	}

}
