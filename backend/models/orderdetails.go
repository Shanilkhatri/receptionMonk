package models

import (
	"log"
	"reakgo/utility"
)

type OrderDetails struct {
	Id              int    `json:"id" db:"id"`
	OrderId         int    `json:"orderid" db:"order_id"`
	PhoneNumber     string `json:"phonenumber" db:"phone_number"`
	SipServer       string `json:"sipserver" db:"sip_server"`
	SipUsername     string `json:"sipusername" db:"sip_username"`
	SipPassword     string `json:"sippassword" db:"sip_password"`
	SipPort         string `json:"sipport" db:"sip_port"`
	IsIvrEnabled    bool   `json:"isivrenabled" db:"is_ivr_enabled"`
	IvrFlow         string `json:"ivrflow" db:"ivr_flow"`
	MaxAllowedUsers int    `json:"maxallowedusers" db:"max_allowed_users"`
	MaxAllowedDepts int    `json:"maxalloweddepts" db:"max_allowed_depts"`
}

type OrderDetailsCondition struct {
	OrderDetails
	CompanyId      int
	AccountType    string
	WhereCondition string
}

func (OrderDetails) OrderDetailsGet(filter OrderDetailsCondition) ([]OrderDetails, error) {
	var orderDetail []OrderDetails
	query := "SELECT * form OrderDetails Where 1=1" + filter.WhereCondition
	condi := map[string]interface{}{
		"Id":        filter.Id,
		"CompanyId": filter.CompanyId,
	}
	rows, err := utility.Db.NamedQuery(query, condi)
	if err != nil {
		log.Println(err)
		return orderDetail, err
	}
	defer rows.Close()
	for rows.Next() {
		var singleRow OrderDetails
		err := rows.Scan(&singleRow.Id, &singleRow.OrderId, &singleRow.PhoneNumber, &singleRow.SipServer, &singleRow.SipUsername, &singleRow.SipPassword, &singleRow.SipPort, &singleRow.IsIvrEnabled, &singleRow.IvrFlow, &singleRow.MaxAllowedUsers, &singleRow.MaxAllowedDepts)
		if err != nil {
			log.Println(err)
			return orderDetail, err
		}
		orderDetail = append(orderDetail, singleRow)
	}
	return orderDetail, err
}

func (OrderDetails) GetFilterOrderDetails(para OrderDetailsCondition) OrderDetailsCondition {
	if para.Id != 0 {
		para.WhereCondition += " AND id=:Id "
	}
	if para.CompanyId != 0 {
		para.WhereCondition += " AND companyid=:CompanyId "
	}

	if para.AccountType != "" {
		if para.AccountType == "owner" {
			para.WhereCondition += " AND type IN ('owner')"
		}
	}

	return para
}
