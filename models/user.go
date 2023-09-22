package models

import (
	"github.com/jmoiron/sqlx"
)

type Users struct {
	ID                    int    `json:"id" db:"id"`
	Name                  string `json:"name" db:"name"`
	Email                 string `json:"email" db:"email"`
	PasswordHash          string `json:"passwordHash" db:"passwordHash"`
	TwoFactorKey          string `json:"twoFactorKey" db:"twoFactorKey"`
	TwoFactorRecoveryCode string `json:"twoFactorRecoveryCode" db:"twoFactorRecoveryCode"`
	DOB                   string `json:"dob" db:"dob"`
	AccountType           string `json:"accountType" db:"accountType"`
	CompanyID             int    `json:"companyId" db:"companyId"`
	Status                string `json:"status" db:"status"`
}

type UserModel struct {
	DB *sqlx.DB
}

func (UserModel) PutUser(tablename string, structure Users) error {
	// sql INSERTION with ORM here
	return nil
}

func (UserModel) PostUser(tablename string, structure Users) error {
	// sql Updation with ORM here
	return nil
}
