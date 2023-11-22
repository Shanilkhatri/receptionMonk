package models

import (
	"fmt"
	"log"
	"reakgo/utility"

	"github.com/jmoiron/sqlx"
)

type Users struct {
	ID                    int64  `json:"id" db:"id"`
	Name                  string `json:"name" db:"name"`
	Email                 string `json:"email" db:"email"`
	PasswordHash          string `json:"passwordHash" db:"passwordHash"`
	TwoFactorKey          string `json:"twoFactorKey" db:"twoFactorKey"`
	TwoFactorRecoveryCode string `json:"twoFactorRecoveryCode" db:"twoFactorRecoveryCode"`
	DOB                   string `json:"dob" db:"dob"`
	AccountType           string `json:"accountType" db:"accountType"`
	CompanyID             int64  `json:"companyId" db:"companyId"`
	Status                string `json:"status" db:"status"`
}

type UserCondition struct {
	Users
	WhereCondition string
}

// for ORM
// func (user Users) PutUser(tablename string, structure Users) error {
// 	// sql INSERTION with ORM here
// 	return nil
// }

// func (user Users) PostUser(tablename string, structure Users) error {
// 	// sql Updation with ORM here
// 	return nil
// }

// insert data in authentication table
func (user Users) PutUser(add Users, tx *sqlx.Tx) bool {

	_, err := tx.NamedExec("INSERT INTO `authentication` (name,email,passwordHash,twoFactorKey,twoFactorRecoveryCode,dob,accountType,companyId,status) VALUES ( :Name,:Email,:PasswordHash,:TwoFactorKey,:TwoFactorRecoveryCode,:DOB,:AccountType,:CompanyID,:Status)", map[string]interface{}{"Name": add.Name, "Email": add.Email, "PasswordHash": add.PasswordHash, "TwoFactorKey": add.TwoFactorKey, "TwoFactorRecoveryCode": add.TwoFactorRecoveryCode, "DOB": add.DOB, "AccountType": add.AccountType, "CompanyID": add.CompanyID, "Status": add.Status})
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

// update user records by id
func (user Users) PostUser(usr Users) (bool, error) {
	userData, err := utility.Db.NamedExec("UPDATE `authentication` SET name=:Name,email=:Email,passwordHash=:PasswordHash,twoFactorKey=:TwoFactorKey,twoFactorRecoveryCode=:TwoFactorRecoveryCode,dob =:DOB,accountType=:AccountType,companyId=:CompanyID, status=:Status WHERE id=:ID ", map[string]interface{}{"Name": usr.Name, "Email": usr.Email, "PasswordHash": usr.PasswordHash, "TwoFactorKey": usr.TwoFactorKey, "TwoFactorRecoveryCode": usr.TwoFactorRecoveryCode, "DOB": usr.DOB, "AccountType": usr.AccountType, "CompanyID": usr.CompanyID, "Status": usr.Status, "ID": usr.ID})
	// Check error
	if err != nil {
		log.Println("error: ", err)
		// utility.Logger(err)
	} else {
		Rowefffect, _ := userData.RowsAffected()
		if Rowefffect == 0 {
			log.Println("input value is not change with previous one or id= " + fmt.Sprint(usr.ID) + "is not valid")
		}
		return Rowefffect > 0, err
	}
	return false, err
}

func (user Users) GetUserById(userId int64) (Users, error) {
	var selectedRow Users

	err := utility.Db.Get(&selectedRow, "SELECT * FROM authentication WHERE id = ?", userId)

	return selectedRow, err
}

func (Users) GetUser(filter UserCondition) ([]Users, error) {
	var userData []Users
	query := "SELECT id,name,email,dob,account_type,company_id,status from authentication Where 1=1" + filter.WhereCondition
	condition := map[string]interface{}{
		"Id":        filter.ID,
		"Dob":       filter.DOB,
		"CompanyId": filter.CompanyID,
	}
	rows, err := utility.Db.NamedQuery(query, condition)
	if err != nil {
		log.Println(err)
		return userData, err
	}
	defer rows.Close()
	for rows.Next() {
		var singleRow Users
		err := rows.Scan(&singleRow.ID, &singleRow.Name, &singleRow.Email, &singleRow.DOB, &singleRow.AccountType, &singleRow.CompanyID, &singleRow.Status)
		if err != nil {
			log.Println(err)
			return userData, err
		}
		userData = append(userData, singleRow)
	}
	return userData, err
}

func (Users) GetParaForFilterUser(para UserCondition) UserCondition {
	if para.ID != 0 {
		para.WhereCondition += " AND id=:Id "
	}
	if para.CompanyID != 0 {
		para.WhereCondition += " AND companyid=:CompanyId "
	}
	if para.DOB != " " {
		para.WhereCondition += " AND DATE_FORMAT(FROM_UNIXTIME(Dob), '%m%d') = DATE_FORMAT(FROM_UNIXTIME(:Dob),'%m%d')"
	}
	if para.AccountType != "" {
		if para.AccountType == "owner" {
			para.WhereCondition += " AND type IN ('user','owner')"
		} else if para.AccountType == "admin" {
			para.WhereCondition += " AND type IN ('user','owner','admin')"
		} else {
			para.WhereCondition += " AND type IN ('user')"
		}
	}

	return para
}
func (Users) DeleteUser(id int) (bool, error) {
	row, err := utility.Db.Exec("DELETE FROM authentication WHERE id = ?", id)
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
