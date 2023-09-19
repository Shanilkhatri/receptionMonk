package models

import (
	"crypto/rand"
	"log"
	"math/big"
	"reakgo/utility"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	Id             int32
	Email          string
	Password       string
	Token          string
	TokenTimestamp int64
}

type AuthenticationModel struct {
	DB *sqlx.DB
}

func (auth AuthenticationModel) GetUserByEmail(email string) (Authentication, error) {
	var selectedRow Authentication

	rows := utility.Db.QueryRow("SELECT * FROM authentication WHERE email = ?", email)
	err := rows.Scan(&selectedRow.Id, &selectedRow.Email, &selectedRow.Password, &selectedRow.Token, &selectedRow.TokenTimestamp)
	if err != nil {
		log.Println("Couldn't lookup user with email : " + email)
		log.Println(err)
	}

	return selectedRow, err
}

func (auth AuthenticationModel) ForgotPassword(id int32) (string, error) {
	Token, err := GenerateRandomString(60)
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

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func (auth AuthenticationModel) TokenVerify(token string, newPassword string) (bool, error) {
	var selectedRow Authentication

	rows := utility.Db.QueryRow("SELECT * FROM authentication WHERE token = ?", token)
	err := rows.Scan(&selectedRow.Id, &selectedRow.Email, &selectedRow.Password, &selectedRow.Token, &selectedRow.TokenTimestamp)
	if err != nil {
		log.Println(err)
		return true, err
	}
	if (selectedRow.TokenTimestamp + 360000) > time.Now().Unix() {
		_, err := auth.ChangePassword(newPassword, selectedRow.Id)
		if err != nil {
			return true, err
		} else {
			return false, err
		}
	}
	return false, err
}

func (auth AuthenticationModel) ChangePassword(newPassword string, id int32) (bool, error) {
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
