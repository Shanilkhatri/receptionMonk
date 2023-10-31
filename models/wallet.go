package models

import (
	"fmt"
	"log"
	"reakgo/utility"

	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	Id        int64  `json:"id" db:"id"`
	Charge    string `json:"charge" db:"charge"`
	Reason    string `json:"reason" db:"reason"`
	Cost      int64  `json:"cost" db:"cost"`
	Epoch     int64  `json:"epoch" db:"epoch"`
	CompanyId int64  `json:"companyId" db:"companyId"`
}

func (wallet Wallet) PutWallet(data Wallet, tx *sqlx.Tx) (bool, error) {
	_, err := tx.NamedExec("INSERT INTO `wallet` (charge,reason,cost,epoch,comapanyId) VALUES (:Charge,:Reason,:Cost,:Epoch,:CompanyId)", map[string]interface{}{"Charge": data.Charge, "Reason": data.Reason, "Cost": data.Cost, "Epoch": data.Epoch, "CompanyId": data.CompanyId})
	if err != nil {
		log.Println(err)
		return true, err
	}
	return false, nil
}

func (Wallet) GetWalletById(walletId int64) (Wallet, error) {
	var selectedRow Wallet

	err := utility.Db.Get(&selectedRow, "SELECT * FROM `wallet` WHERE id = ?", walletId)

	return selectedRow, err
}

func (Wallet) PostWallet(updatedWallet Wallet) (bool, error) {
	walletData, err := utility.Db.NamedExec("UPDATE `wallet` SET charge=:Charge,reason=:Reason,cost=:Cost,epoch=:Epoch,companyId=:CompanyId WHERE id=:Id ", map[string]interface{}{"Charge": updatedWallet.Charge, "Reason": updatedWallet.Reason, "Cost": updatedWallet.Cost, "Epoch": updatedWallet.Epoch, "CompanyId": updatedWallet.CompanyId, "Id": updatedWallet.Id})
	// Check error
	if err != nil {
		log.Println("error: ", err)
		utility.Logger(err)
	} else {
		Rowefffect, _ := walletData.RowsAffected()
		if Rowefffect == 0 {
			log.Println("input value is not change with previous one or id= " + fmt.Sprint(updatedWallet.Id) + "is not valid")
		}
		return Rowefffect > 0, err
	}
	return false, err
}
