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

type TicketsCondition struct {
	WhereCondition string
	Tickets
}

func (Tickets) PutTicket(ticketStruct Tickets, tx *sqlx.Tx) bool {

	// executing query
	_, err := tx.NamedExec("INSERT INTO `tickets` (userId,email,customerName,createdTime,lastUpdatedOn,status,query,feedback,lastResponse) VALUES ( :UserId,:Email,:CustomerName,:CreatedTime,:LastUpdatedOn,:Status,:Query,:FeedBack,:LastResponse)", map[string]interface{}{"UserId": ticketStruct.UserId, "Email": ticketStruct.Email, "CustomerName": ticketStruct.CustomerName, "CreatedTime": ticketStruct.CreatedTime, "LastUpdatedOn": ticketStruct.LastUpdatedOn, "Status": ticketStruct.Status, "Query": ticketStruct.Query, "FeedBack": ticketStruct.FeedBack, "LastResponse": ticketStruct.LastResponse, "CompanyId": ticketStruct.CompanyId})
	// Check error
	if err != nil {
		log.Println(err)
		//logger remove for duplicate entry that means duplicate error message not send at email.
		istrue, _ := Helper.CheckSqlError(err, "Duplicate entry")
		if !istrue {
			Helper.Logger(err)
		}
		return false
	} else {
		return true
	}
}

func (Tickets) GetTicketById(ticketId int64) (Tickets, error) {
	var selectedRow Tickets

	err := utility.Db.Get(&selectedRow, "SELECT * FROM `tickets` WHERE id = ?", ticketId)

	return selectedRow, err
}

func (Tickets) GetTicketsByUserId(userId int64) ([]Tickets, error) {
	var tickets []Tickets

	query := "SELECT * FROM `tickets` WHERE 1=1 AND userId=:UserId AND status IN (:Status1, :Status2) "
	condition := map[string]interface{}{
		"UserId":  userId,
		"Status1": "in_process", // only tickets under Process will show up here
		"Status2": "open",       // only tickets under Process will show up here
	}
	rows, err := utility.Db.NamedQuery(query, condition)
	if err != nil {
		log.Println(err)
		return tickets, err
	}
	defer rows.Close()
	for rows.Next() {
		var singleRow Tickets
		err := rows.Scan(&singleRow.Id, &singleRow.UserId, &singleRow.Email, &singleRow.CustomerName, &singleRow.CreatedTime, &singleRow.LastUpdatedOn, &singleRow.Status, &singleRow.Query, &singleRow.FeedBack, &singleRow.LastResponse, &singleRow.CompanyId)
		if err != nil {
			log.Println(err)
			return tickets, err
		}
		tickets = append(tickets, singleRow)
	}
	return tickets, err
}

// update user records by id
func (Tickets) PostTicket(ticketStruct Tickets, tx *sqlx.Tx) (bool, error) {
	userData, err := tx.NamedExec("UPDATE tickets SET userId=:UserId,email=:Email,customerName=:CustomerName,createdTime =:CreatedTime,lastUpdatedOn=:LastUpdatedOn,status=:Status,query=:Query,feedback=:FeedBack, lastResponse=:LastResponse,companyId=:CompanyId WHERE id=:Id ", map[string]interface{}{"UserId": ticketStruct.UserId, "Email": ticketStruct.Email, "CustomerName": ticketStruct.CustomerName, "CreatedTime": ticketStruct.CreatedTime, "LastUpdatedOn": ticketStruct.LastUpdatedOn, "Status": ticketStruct.Status, "Query": ticketStruct.Query, "FeedBack": ticketStruct.FeedBack, "LastResponse": ticketStruct.LastResponse, "CompanyId": ticketStruct.CompanyId, "Id": ticketStruct.Id})
	// Check error
	if err != nil {
		log.Println("error: ", err)
		// utility.Logger(err)
	} else {
		Rowefffect, _ := userData.RowsAffected()
		if Rowefffect == 0 {
			log.Println("input value is not change with previous one or id= " + fmt.Sprint(ticketStruct.Id) + "is not valid")
		}
		return Rowefffect > 0, err
	}
	return false, err
}

func (ord Orders) GetTickets(filter TicketsCondition, tx *sqlx.Tx) ([]Tickets, error) {
	var ticketsArr []Tickets
	query := "SELECT * FROM `tickets` WHERE 1=1 " + filter.WhereCondition + " ORDER BY tickets.id DESC;"
	condtion := map[string]interface{}{
		"id":            filter.Id,
		"createdTime":   filter.CreatedTime,
		"lastUpdatedOn": filter.LastUpdatedOn,
		"companyId":     filter.CompanyId,
		"userId":        filter.UserId,
	}
	rows, err := tx.NamedQuery(query, condtion)
	if err != nil {
		return ticketsArr, err
	}
	defer rows.Close()
	for rows.Next() {
		var singleRow Tickets
		err := rows.Scan(&singleRow.Id, &singleRow.UserId, &singleRow.Email, &singleRow.CustomerName, &singleRow.CreatedTime, &singleRow.LastUpdatedOn, &singleRow.Status, &singleRow.Query, &singleRow.FeedBack, &singleRow.LastResponse, &singleRow.CompanyId)
		if err != nil {
			return ticketsArr, err
		}

		ticketsArr = append(ticketsArr, singleRow)
	}
	return ticketsArr, nil
}

// Building conditions for getting Tickets Data
func (Tickets) GetParamsForFilterTicketsData(params TicketsCondition) TicketsCondition {

	if params.Id != 0 {
		params.WhereCondition += " AND tickets.id = :id"
	}
	if params.UserId != 0 {
		params.WhereCondition += " AND tickets.userId = :userId"
	}
	if params.LastUpdatedOn != 0 {
		params.WhereCondition += " AND tickets.lastUpdatedOn = :lastUpdatedOn"
	} else if params.CreatedTime != 0 {
		params.WhereCondition += " AND tickets.createdTime = :createdTime"
	} else if params.CompanyId != 0 {
		params.WhereCondition += " AND tickets.companyId= :companyId"
	}
	return params
}

func (Tickets) DeleteTicket(id int) (bool, error) {
	row, err := utility.Db.Exec("UPDATE tickets SET status=:Status WHERE id = :Id", map[string]interface{}{"Status": "archive", "Id": id})
	// row, err := utility.Db.Exec("DELETE FROM tickets WHERE id = ?", id)
	if err != nil {
		log.Print(err)
		return false, err
	}
	rowsDeleted, err := row.RowsAffected()
	if err != nil {
		log.Print(err)
		return false, err
	}
	return rowsDeleted > 0, nil
}
