package models

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

type Extensions struct {
	Id          int    `json:"id" db:"id" primarykey:"true"`
	Extension   string `json:"extension" db:"extension"`
	UserId      int    `json:"userid" db:"user_id"`
	Department  int    `json:"department" db:"department"`
	SipServer   string `json:"sipserver" db:"sip_server"`
	SipUserName string `json:"sipusername" db:"sip_username"`
	SipPassword string `json:"sippassword" db:"sip_password"`
	SipPort     string `json:"sipport" db:"sip_port"`
	db          Database
}

func (extension Extensions) PostExtension(data Extensions, tx *sqlx.Tx) (bool, error) {
	row, err := tx.NamedExec("UPDATE `extensions` INNER JOIN `authentication` ON `extensions`.`user_id` = `authentication`.`id` INNER JOIN `company` ON `authentication`.`company_id` = `company`.`id` SET extension=:Extension,department=:Department,sip_server=:SipServer,sip_username=:SipUserName,sip_password:=SipPassword,sip_port:=SipPort WHERE `extensions`.`id`=:Id ", map[string]interface{}{"Extension": data.Extension, "Department": data.Department, "SipServer": data.SipServer, "SipUserName": data.SipUserName, "SipPassword": data.SipPassword, "SipPort": data.SipPort, "Id": data.Id})
	if err != nil {
		log.Println(err)
		return false, err
	}
	rowUpdate, _ := row.RowsAffected()
	return rowUpdate > 0, nil
}

type Users struct {
	Id                    int    `json:"id" db:"id"`
	Name                  string `json:"name" db:"name"`
	Email                 string `json:"email" db:"email"`
	PasswordHash          string `json:"passwordHash" db:"passwordHash"`
	TwoFactorKey          string `json:"twoFactorKey" db:"twoFactorKey"`
	TwoFactorRecoveryCode string `json:"twoFactorRecoveryCode" db:"twoFactorRecoveryCode"`
	DOB                   string `json:"dob" db:"dob"`
	AccountType           string `json:"accountType" db:"accountType"`
	CompanyId             int    `json:"companyId" db:"companyId"`
	Status                string `json:"status" db:"status"`
	Token                 string `json:"token" db:"token"`
}

// Database interface to abstract database operations
type Database interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
}

// SQLxDB is a mock implementation of the sqlx.DB interface
type MockDB struct {
	mock.Mock
}

func (m *MockDB) NamedExec(query string, arg interface{}) (sql.Result, error) {
	args := m.Called(query, arg)
	return args.Get(0).(sql.Result), args.Error(1)
}

type Exten struct {
	db Database
}

func TestPostExtensions(t *testing.T) {
	// Create a mock database
	mockDB := new(MockDB)

	// Create an instance of Exten with the mock database
	exten := Extensions{db: mockDB}

	// Define the expected query and arguments
	expectedQuery := "UPDATE `extensions` INNER JOIN `authentication` ON `extensions`.`user_id` = `authentication`.`id` INNER JOIN `company` ON `authentication`.`company_id` = `company`.`id` SET extension=:Extension,department=:Department,sip_server=:SipServer,sip_username=:SipUserName,sip_password:=SipPassword,sip_port:=SipPort WHERE `extensions`.`id`=:Id "
	expectedArgs := map[string]interface{}{
		"Extension":  "mockExtension",
		"Department": "mockDepartment",
		"SipSe":      "mockSipSe",
		"SipUser":    "mockSipUser",
		"SipPass":    "mockSipPass",
		"SipPort":    "mockSipPort",
		"Id":         123,
	}

	// Set up expectations for the mock database
	mockDB.On("NamedExec", expectedQuery, expectedArgs).Return(sqlmock.NewResult(1, 1), nil)

	// Call the function under test
	result, err := exten.PostExtension(Extensions{}, nil)

	// Check if the function behaves as expected
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if !result {
		t.Error("Expected result to be true, but got false")
	}

	// Assert that the expected database operation was called
	mockDB.AssertExpectations(t)
}
