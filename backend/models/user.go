package models

import (
	"log"
	"reakgo/utility"
)

type User struct {
	Id                    int64  `json:"id" db:"id"`
	Name                  string `json:"name" db:"name"`
	Email                 string `json:"email" db:"email"`
	PasswordHash          string `json:"passwordHash" db:"password_hash"`
	TwoFactorKey          string `json:"twoFactorKey" db:"two_factor_key"`
	TwoFactorRecoveryCode string `json:"twoFactorRecoveryCode" db:"two_factor_recovery_code"`
	Dob                   string `json:"dob" db:"dob"`
	AuthToken             string `json:"authToken" db:"auth_token"`
	AccountType           string `json:"accountType" db:"account_type"`
	CompanyId             int64  `json:"companyId" db:"company_id"`
	Status                string `json:"status" db:"status"`
}

type UserCondition struct {
	User
	WhereCondition string
}

func (User) GetUser(filter UserCondition) ([]User, error) {
	var userData []User
	query := "SELECT id,name,email,dob,account_type,company_id,status form authentication Where 1=1" + filter.WhereCondition
	condi := map[string]interface{}{
		"Id":        filter.Id,
		"Dob":       filter.Dob,
		"CompanyId": filter.CompanyId,
	}
	rows, err := utility.Db.NamedQuery(query, condi)
	if err != nil {
		log.Println(err)
		return userData, err
	}
	defer rows.Close()
	for rows.Next() {
		var singleRow User
		err := rows.Scan(&singleRow.Id, &singleRow.Name, &singleRow.Email, &singleRow.Dob, &singleRow.AccountType, &singleRow.CompanyId, &singleRow.Status)
		if err != nil {
			log.Println(err)
			return userData, err
		}
		userData = append(userData, singleRow)
	}
	return userData, err
}

func (User) GetParaForFilterUser(para UserCondition) UserCondition {
	if para.Id != 0 {
		para.WhereCondition += " AND id=:Id "
	}
	if para.CompanyId != 0 {
		para.WhereCondition += " AND companyid=:CompanyId "
	}
	if para.Dob != " " {
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

func (User) DeleteUser(id int) (bool, error) {
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
