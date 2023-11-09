package models

import (
	"database/sql"
	"fmt"
	"log"
	"reakgo/utility"

	"time"

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
	Token                 string
	TokenTimestamp        int64 `db:"tokenTimestamp"`
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

func (auth Authentication) TokenVerify(token string, newPassword string) (bool, error) {
	var selectedRow Authentication

	rows := utility.Db.QueryRow("SELECT * FROM authentication WHERE token = ?", token)
	err := rows.Scan(&selectedRow.ID, &selectedRow.Email, &selectedRow.PasswordHash, &selectedRow.Token, &selectedRow.TokenTimestamp)
	if err != nil {
		log.Println(err)
		return true, err
	}
	if (selectedRow.TokenTimestamp + 360000) > time.Now().Unix() {
		_, err := auth.ChangePassword(newPassword, int32(selectedRow.ID))
		if err != nil {
			return true, err
		} else {
			return false, err
		}
	}
	return false, err
}

func (auth Authentication) ChangePassword(newPassword string, id int32) (bool, error) {
	query, err := utility.Db.Prepare("UPDATE authentication SET password = ? WHERE id = ?")
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
	Email        string `json:"authEmailId"`
	Password     string `json:"password"`
	Otp          string `json:"authSignInOTP"`
	EmailToken   string `json:"emailVerToken"`
	EpochCurrent int64  `json:"epochcurrent"`
	EpochExpired int64
}

func (user Authentication) GetUserByEmailIds(data SignupDetails) (bool, error) {

	var selectedRow SignupDetails
	if EmailExistOrNot(data.Email) {
		//update data
		queryOfauth, err := utility.Db.NamedExec("UPDATE `authentication` SET `emailToken`=:EmailToken,otp=:Otp,epoch=:Epoch WHERE email=:Email ", map[string]interface{}{"EmailToken": data.EmailToken, "OTP": data.Otp, "EpochCurrent": data.EpochCurrent, "EpochExpired": data.EpochExpired, "Email": data.Email})
		// Check error
		if err != nil {
			log.Println(err)
			return false, err
		}
		Rowefffect, _ := queryOfauth.RowsAffected()
		return Rowefffect > 0, err

	} else {
		_, err := utility.Db.NamedExec("INSERT INTO `authentication` (emailToken,otp,epoch,email) VALUES (:EmailToken,:Otp,:Epoch,:Email)", map[string]interface{}{"EmailToken": selectedRow.EmailToken, "Otp": selectedRow.Otp, "EpochCurrent": selectedRow.EpochCurrent, "EpochExpired": selectedRow.EpochExpired, "Email": selectedRow.Email})
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

	err := utility.Db.Get(&selectedRow, "SELECT currenttime,expiredtime,otp FROM authentication WHERE email = ?", email)

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
