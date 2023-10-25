package models

import (
	"fmt"
	"log"
	"reakgo/utility"

	"github.com/jmoiron/sqlx"
)

// - id                int64           `UNIQUE/PRIMARY`
// - email      	   string
// - userId      	   int64
// - customerName     string
// - created_time      int64           `EPOCH-TIMESTAMP`
// - last_updated_on   int64           `EPOCH-TIMESTAMP`
// - status            string          `ENUM(open,in_process,closed)`
// - query             string
// - feedback          string          `ENUM (satisfied,not_satisfied,no_feedback)`
// - last_response     string

type Tickets struct {
	Id            int64  `db:"id" json:"id"`
	UserId        int64  `db:"userId" json:"userId"`
	Email         string `db:"email" json:"email"`
	CustomerName  string `db:"customerName" json:"customerName"`
	CreatedTime   int64  `db:"createdTime" json:"createdTime"`
	LastUpdatedOn int64  `db:"lastUpdatedOn" json:"lastUpdatedOn"`
	Status        string `db:"status" json:"status"`
	Query         string `db:"query" json:"query"`
	FeedBack      string `db:"feedback"`
	LastResponse  string `db:"lastResponse" json:"lastResponse"`
	CompanyId     int64  `db:"companyId" json:"companyId"`
}

func (Tickets) PutTicket(ticketStruct Tickets, tx *sqlx.Tx) bool {

	// executing query
	_, err := tx.NamedExec("INSERT INTO `tickets` (userId,email,customerName,createdTime,lastUpdatedOn,status,query,feedback,lastResponse) VALUES ( :UserId,:Email,:CustomerName,:CreatedTime,:LastUpdatedOn,:Status,:Query,:FeedBack,:LastResponse)", map[string]interface{}{"UserId": ticketStruct.UserId, "Email": ticketStruct.Email, "CustomerName": ticketStruct.CustomerName, "CreatedTime": ticketStruct.CreatedTime, "LastUpdatedOn": ticketStruct.LastUpdatedOn, "Status": ticketStruct.Status, "Query": ticketStruct.Query, "FeedBack": ticketStruct.FeedBack, "LastResponse": ticketStruct.LastResponse, "CompanyId": ticketStruct.CompanyId})
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

func (Tickets) GetTicketById(ticketId int64) (Tickets, error) {
	var selectedRow Tickets

	err := utility.Db.Get(&selectedRow, "SELECT * FROM tickets WHERE id = ?", ticketId)

	return selectedRow, err
}

// update user records by id
func (Tickets) PostTicket(ticketStruct Tickets) (bool, error) {
	userData, err := utility.Db.NamedExec("UPDATE tickets SET userId=:UserId,email=:Email,customerName=:CustomerName,createdTime =:CreatedTime,lastUpdatedOn=:LastUpdatedOn,status=:Status,query=:Query,feedback=:FeedBack, lastResponse=:LastResponse,companyId=:CompanyId WHERE id=:Id ", map[string]interface{}{"UserId": ticketStruct.UserId, "Email": ticketStruct.Email, "CustomerName": ticketStruct.CustomerName, "CreatedTime": ticketStruct.CreatedTime, "LastUpdatedOn": ticketStruct.LastUpdatedOn, "Status": ticketStruct.Status, "Query": ticketStruct.Query, "FeedBack": ticketStruct.FeedBack, "LastResponse": ticketStruct.LastResponse, "CompanyId": ticketStruct.CompanyId, "Id": ticketStruct.Id})
	// Check error
	if err != nil {
		utility.Logger(err)

	} else {
		Rowefffect, _ := userData.RowsAffected()
		if Rowefffect == 0 {
			log.Println("input value is not change with previous one or id= " + fmt.Sprint(ticketStruct.Id) + "is not valid")
		}
		return Rowefffect > 0, err
	}
	return false, err
}
