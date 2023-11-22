package models

import (
	"database/sql"
	"fmt"
	"log"
	"reakgo/utility"

	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

//	type Authentication struct {
//		Id             int32
//		Email          string
//		Password       string
//		Token          string
//		TokenTimestamp int64 `db:"tokenTimestamp"`
//	}
type Authentication struct {
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
	Token                 string `json:"token" db:"token"`
	EmailToken            string `json:"emailToken" db:"emailToken"`
	Otp                   string `json:"otp" db:"otp"`
	EpochCurrent          int64  `db:"epochcurrent"`
	EpochExpired          int64  `db:"epochexpired"`
	IsWizardComplete      string `json:"iswizardcomplete" db:"iswizardcomplete"`
}

type TwoFactor struct {
	UserId int32  `db:"userId"`
	Secret string `db:"secret"`
}

func (auth Authentication) GetUserByEmail(email string) (Authentication, error) {
	var selectedRow Authentication

	err := utility.Db.Get(&selectedRow, "SELECT * FROM authentication WHERE email = ?", email)

	return selectedRow, err
}

func (auth Authentication) ForgotPassword(id int32) (string, error) {
	Token, err := utility.GenerateRandomString(60)
	if err != nil {
		log.Println("Random String Generator Failed")
	}
	TokenTimestamp := time.Now().Unix()
	query, err := utility.Db.Prepare("UPDATE authentication SET Token = ?, TokenTimestamp = ? WHERE id = ?")
	if err != nil {
		log.Println("MySQL Query Failed")
	}
	_, err = query.Exec(Token, TokenTimestamp, id)
	if err != nil {
		log.Println(err)
	} else {
		// pass
	}
	return Token, err
}

// func (auth Authentication) TokenVerify(token string, newPassword string) (bool, error) {
// 	var selectedRow Authentication

// 	rows := utility.Db.QueryRow("SELECT * FROM authentication WHERE token = ?", token)
// 	err := rows.Scan(&selectedRow.ID, &selectedRow.Email, &selectedRow.PasswordHash, &selectedRow.Token, &selectedRow.TokenTimestamp)
// 	if err != nil {
// 		log.Println(err)
// 		return true, err
// 	}
// 	if (selectedRow.TokenTimestamp + 360000) > time.Now().Unix() {
// 		_, err := auth.ChangePassword(newPassword, int32(selectedRow.ID))
// 		if err != nil {
// 			return true, err
// 		} else {
// 			return false, err
// 		}
// 	}
// 	return false, err
// }

func (auth Authentication) ChangePassword(newPassword string, id int32) (bool, error) {
	query, err := utility.Db.Prepare("UPDATE authentication SET passwordHash = ? WHERE id = ?")
	if err != nil {
		log.Println("MySQL Query Failed")
	}
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	if err != nil {
		log.Println(err)
		return true, err
	}
	_, err = query.Exec(newPasswordHash, id)
	if err != nil {
		log.Println(err)
		return true, err
	} else {
		return false, err
	}
}

func (auth Authentication) TwoFactorAuthAdd(secret string, userId int) (bool, error) {
	_, err := utility.Db.NamedExec("INSERT INTO twoFactor (userId, secret) VALUES(:id, :2faSecret) ON DUPLICATE KEY UPDATE secret=:2faSecret", map[string]interface{}{"2faSecret": secret, "id": userId})
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func (auth Authentication) CheckTwoFactorRegistration(userId int32) string {
	twoFactor := TwoFactor{}
	utility.Db.Get(&twoFactor, "SELECT * FROM twoFactor WHERE userId = ?", userId)
	return twoFactor.Secret
}

func (auth Authentication) GetAllAuthRecords() ([]Authentication, error) {
	var allAuthenticationRows []Authentication

	// SQL query to select all rows from the Abc table
	query := "SELECT * FROM authentication"

	// Execute the query and scan the results into the Abc slice
	err := utility.Db.Select(&allAuthenticationRows, query)
	if err != nil {
		return nil, err
	}

	return allAuthenticationRows, nil
}

func (auth Authentication) GetAuthenticationByToken(token string) (*Authentication, error) {

	var authO Authentication

	// SQL query to select a row by Token
	query := "SELECT * FROM authentication WHERE Token = ? LIMIT 1"

	// Execute the query and scan the result into the Authentication struct
	err := utility.Db.Get(&authO, query, token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Authentication not found for Token: %s", token)
		}
		return nil, err
	}
	return &authO, nil
}

type SignupDetails struct {
	Email        string `json:"authEmailId" db:"email"`
	PasswordHash string `json:"passwordHash" db:"passwordHash"`
	Otp          string `json:"otp" db:"otp"`
	EmailToken   string `json:"emailVerToken" db:"emailToken"`
	EpochCurrent int64  `json:"epochcurrent" db:"epochcurrent"`
	EpochExpired int64  `db:"epochexpired"`
	Token        string `json:"token" db:"token"`
}

func (user Authentication) GetUserByEmailIds(data SignupDetails) (bool, error) {
	log.Println("data", data)
	log.Println("email exists: ", EmailExistOrNot(data.Email))
	if EmailExistOrNot(data.Email) {
		log.Println("hi update")
		//update data
		queryOfauth, err := utility.Db.NamedExec("UPDATE `authentication` SET `emailToken`=:EmailToken,`otp`=:Otp,`epochcurrent`=:EpochCurrent,`epochexpired`=:EpochExpired, `token`=:Token WHERE `email`=:Email ", map[string]interface{}{"EmailToken": data.EmailToken, "Otp": data.Otp, "EpochCurrent": data.EpochCurrent, "EpochExpired": data.EpochExpired, "Email": data.Email, "Token": data.Token})
		// Check error
		if err != nil {
			log.Println(err)
			return false, err
		}
		Rowefffect, _ := queryOfauth.RowsAffected()
		return Rowefffect > 0, err
	} else {
		log.Println("insert :", data)
		_, err := utility.Db.NamedExec("INSERT INTO `authentication` (emailToken,otp,epochcurrent,epochexpired,email,passwordHash,token) VALUES (:EmailToken,:Otp,:EpochCurrent,:EpochExpired,:Email,:PasswordHash,:Token)", map[string]interface{}{"EmailToken": data.EmailToken, "Otp": data.Otp, "EpochCurrent": data.EpochCurrent, "EpochExpired": data.EpochExpired, "Email": data.Email, "PasswordHash": data.PasswordHash, "Token": data.Token})
		// Check error
		if err != nil {
			log.Println(err)
			return false, err
		} else {
			return true, nil
		}
	}

}

func EmailExistOrNot(email string) bool {
	var countMatchId int64
	err := utility.Db.QueryRow("SELECT count(*) FROM authentication WHERE email = ?", email).Scan(&countMatchId)
	//check error
	if err != nil {
		log.Println(err)
		return false
	} else {
		return countMatchId > 0
	}
}
func (auth Authentication) GetUserDetailsByEmail(email string) (SignupDetails, error) {
	var selectedRow SignupDetails

	log.Println("email: ", email)
	err := utility.Db.Get(&selectedRow, "SELECT emailToken, otp,epochcurrent,epochexpired,email FROM `authentication` WHERE email = ?", email)
	log.Println("selectedRow: ", selectedRow)
	return selectedRow, err
}

// func (user Authentication) GetUserByEmailIds(data SignupDetails) (Authentication, error) {
// 	var selectedRow Authentication
// 	if EmailExistOrNot(data.Email) {
// 		err := utility.Db.Get(&selectedRow, "SELECT * FROM authentication WHERE email = ?", data.Email)
// 		return selectedRow, err
// 	} else {
// 		_, err := utility.Db.NamedExec("INSERT INTO `authentication` (email) VALUES (:Email)", map[string]interface{}{"Email": selectedRow.Email})
// 		// Check error
// 		if err != nil {
// 			log.Println(err)
// 			return selectedRow, err
// 		} else {
// 			return selectedRow, nil
// 		}
// 	}

// }
func (user Authentication) UpdateCompanyIdByEmail(id int64, companyId int64, tx *sqlx.Tx) (bool, error) {
	//update data
	queryOfauth, err := tx.NamedExec("UPDATE `authentication` SET `companyId`=:CompanyId  WHERE `id`=:Id ", map[string]interface{}{"CompanyId": companyId, "Id": id})
	// Check error
	if err != nil {
		log.Println(err)
		return false, err
	}
	Rowefffect, _ := queryOfauth.RowsAffected()
	return Rowefffect > 0, err
}
