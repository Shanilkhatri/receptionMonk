package models

import (
	"log"
	"reakgo/utility"

	"github.com/jmoiron/sqlx"
)

type Extensions struct {
	Id          int64  `json:"id" db:"id" primarykey:"true"`
	Extension   string `json:"extension" db:"extension"`
	UserId      int64  `json:"userid" db:"user_id"`
	Department  int64  `json:"department" db:"department"`
	SipServer   string `json:"sipserver" db:"sip_server"`
	SipUserName string `json:"sipusername" db:"sip_username"`
	SipPassword string `json:"sippassword" db:"sip_password"`
	SipPort     string `json:"sipport" db:"sip_port"`
}

type ExtensionCondition struct {
	Extensions
	CompanyId      int64
	AccountType    string
	WhereCondition string
}

func (extension Extensions) PostExtension(data Extensions, tx *sqlx.Tx) (bool, error) {
	log.Println("extension Struct: ", data)
	row, err := tx.NamedExec("UPDATE `extension` INNER JOIN `authentication` ON `extension`.`user_id` = `authentication`.`id` INNER JOIN `company` ON `authentication`.`companyId` = `company`.`id` SET extension=:Extension,department=:Department,sip_server=:SipServer,sip_username=:SipUserName,sip_password=:SipPassword,sip_port=:SipPort WHERE `extension`.`id`=:Id ", map[string]interface{}{"Extension": data.Extension, "Department": data.Department, "SipServer": data.SipServer, "SipUserName": data.SipUserName, "SipPassword": data.SipPassword, "SipPort": data.SipPort, "Id": data.Id})
	if err != nil {
		log.Println(err)
		return false, err
	}
	rowUpdate, _ := row.RowsAffected()
	return rowUpdate > 0, nil
}

func (extension Extensions) PutExtension(data Extensions, tx *sqlx.Tx) (bool, error) {
	_, err := tx.NamedExec("INSERT INTO `extension` (extension,user_id,department,sip_server,sip_username,sip_password,sip_port) VALUES (:Extension,:UserId,:Department,:SipServer,:SipUserName,:SipPassword,:SipPort)", map[string]interface{}{"Extension": data.Extension, "UserId": data.UserId, "Department": data.Department, "SipServer": data.SipServer, "SipUserName": data.SipUserName, "SipPassword": data.SipPassword, "SipPort": data.SipPort})
	if err != nil {
		log.Println(err)
		return true, err
	}
	return false, nil
}

func (extension Extensions) DeleteExtensionDetail(id int) (bool, error) {
	row, err := utility.Db.Exec("DELETE FROM extension WHERE id = ?", id)
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
func (extension Extensions) GetExtensionById(extensionId int64) (Extensions, error) {
	var fetchedExtension Extensions
	err := utility.Db.Get(&fetchedExtension, "SELECT * FROM extension WHERE id = ?", extensionId)
	return fetchedExtension, err
}
func (extension Extensions) GetExtensions(filter ExtensionCondition) ([]Extensions, error) {
	var ExtensionData []Extensions
	query := "SELECT * FROM extension Where 1=1" + filter.WhereCondition
	condi := map[string]interface{}{
		"Id":     filter.Id,
		"UserId": filter.UserId,
	}
	rows, err := utility.Db.NamedQuery(query, condi)
	if err != nil {
		log.Println(err)
		return ExtensionData, err
	}
	defer rows.Close()
	for rows.Next() {
		var singleRow Extensions
		err := rows.Scan(&singleRow.Id, &singleRow.Extension, &singleRow.UserId, &singleRow.Department, &singleRow.SipServer, &singleRow.SipUserName, &singleRow.SipPassword, &singleRow.SipPort)
		if err != nil {
			log.Println(err)
			return ExtensionData, err
		}
		ExtensionData = append(ExtensionData, singleRow)
	}
	return ExtensionData, err
}

func (extension Extensions) GetParaForFilterExtension(para ExtensionCondition) ExtensionCondition {
	if para.Id != 0 {
		para.WhereCondition += " AND id=:Id "
	}
	if para.UserId != 0 {
		para.WhereCondition += " AND user_id=:UserId "
	}
	// if para.AccountType != "" {
	// 	if para.AccountType == "owner" {
	// 		para.WhereCondition += " AND type IN ('user','owner')"
	// 	} else if para.AccountType == "admin" {
	// 		para.WhereCondition += " AND type IN ('user','owner','admin')"
	// 	} else {
	// 		para.WhereCondition += " AND type IN ('user')"
	// 	}
	// }

	return para
}
