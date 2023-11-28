package models

import (
	"fmt"
	"log"
	"reakgo/utility"

	"github.com/jmoiron/sqlx"
)

type KycDetails struct {
	Id         int64  `json:"kycId" db:"id"`
	UserId     int64  `json:"userid" db:"userid"`
	DocName    string `json:"doc_name" db:"doc_name"`
	DocPicName string `json:"doc_pic_name" db:"doc_pic_name"`
}

func (KycDetails) Putkyc(add KycDetails, tx *sqlx.Tx) bool {
	_, err := tx.NamedExec("INSERT INTO `kyc_details` (userid,doc_name,doc_pic_name) VALUES (:Userid,:Doc_name,:Doc_pic_name)", map[string]interface{}{"Userid": add.UserId, "Doc_name": add.DocName, "Doc_pic_name": add.DocPicName})
	// Check error
	if err != nil {
		log.Println(err)
		//logger remove for duplicate entry that means duplicate error message not send at email.
		istrue, _ := utility.CheckSqlError(err, "Duplicate entry")
		if !istrue {
			utility.Logger(err)
		}
	} else {
		return true
	}
	return false
}

// update user kyc records by userid
func (KycDetails) Postkyc(usr KycDetails, tx *sqlx.Tx) (bool, error) {
	userData, err := tx.NamedExec("UPDATE `kyc_details` SET doc_name=:Doc_name,doc_pic_name=:Doc_pic_name WHERE userid=:Userid ", map[string]interface{}{"Doc_name": usr.DocName, "Doc_pic_name": usr.DocPicName, "Userid": usr.UserId})
	// Check error
	if err != nil {
		log.Println("error: ", err)
		// utility.Logger(err)
	} else {
		Rowefffect, _ := userData.RowsAffected()
		if Rowefffect == 0 {
			log.Println("input value is not change with previous one or userid= " + fmt.Sprint(usr.UserId) + "is not valid")
		}
		return Rowefffect > 0, err
	}
	return false, err
}
func (KycDetails) Getkyc(userId int64) (KycDetails, error) {
	var selectedRow KycDetails

	err := utility.Db.Get(&selectedRow, "SELECT * FROM kyc_details WHERE userid = ?", userId)

	return selectedRow, err
}
