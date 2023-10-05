package models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Extensions struct {
	Id          int    `json:"id" db:"id" primarykey:"true"`
	Extension   string `json:"extension" db:"extension"`
	UserId      int    `json :"userid" db:"user_id"`
	Department  int    `json:"department" db:"department"`
	SipServer   string `json:"sipserver" db:"sip_server"`
	SipUserName string `json :"sipusername" db:"sip_username"`
	SipPassword string `json :"sippassword" db:"sip_password"`
	SipPort     string `json :"sipport" db:"sip_port"`
}

func (extension Extensions) PostExtension(data Extensions, tx *sqlx.Tx) (bool, error) {
	row, err := tx.NamedExec("UPDATE `extensions` INNER JOIN `authentication` ON `extensions`.`user_id` = `authentication`.`id` INNER JOIN `company` ON `authentication`.`company_id` = `company`.`id` SET extension=:Extension,department=:Department,sip_server=:SipServer,sip_username=:SipUserName,sip_password:=SipPassword,sip_port:=SipPort WHERE id=:Id ", map[string]interface{}{"Extension": data.Extension, "Department": data.Department, "SipServer": data.SipServer, "SipUserName": data.SipUserName, "SipPassword": data.SipPassword, "SipPort": data.SipPort, "Id": data.Id})
	if err != nil {
		log.Println(err)
		return false, err
	}
	rowUpdate, _ := row.RowsAffected()
	return rowUpdate > 0, nil
}
