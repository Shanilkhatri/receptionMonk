package models

import (
	"fmt"
	"log"
	"reakgo/utility"
)

type Products struct {
	Id           int64  `db:"id" json:"productId"`
	Name         string `db:"name" json:"productName"`
	Price        int64  `db:"price" json:"price"`
	PlanValidity int64  `db:"plan_validity" json:"plan_validity"`
	Description  string `db:"description" json:"description"`
	Status       string `db:"status" json:"productStatus" `
}

// get products
func GetProductById(id int64) (Products, error) {
	var selectedRow Products
	err := utility.Db.Get(&selectedRow, "SELECT * FROM products WHERE id = ?", id)
	return selectedRow, err
}
func (Products) PutProduct(add Products) bool {
	_, err := utility.Db.NamedExec("INSERT INTO `products` (name,price,plan_validity,description,status) VALUES (:Name,:Price,:Plan_validity,:Description,:Status)", map[string]interface{}{"Name": add.Name, "Price": add.Price, "Plan_validity": add.PlanValidity, "Description": add.Description, "Status": add.Status})
	// Check error
	if err != nil {
		log.Println(err)
		//logger remove for duplicate entry that means duplicate error message not send at email.
		istrue, _ := Helper.CheckSqlError(err, "Duplicate entry")
		if !istrue {
			Helper.Logger(err, false)
		}
		return false
	} else {
		return true
	}
}
func (Products) PostProduct(usr Products) (bool, error) {
	userData, err := utility.Db.NamedExec("UPDATE `products` SET name=:Name,price=:Price,plan_validity=:Plan_validity,description=:Description,status=:Status WHERE id=:ID ", map[string]interface{}{"Name": usr.Name, "Price": usr.Price, "Plan_validity": usr.PlanValidity, "Description": usr.Description, "Status": usr.Status, "ID": usr.Id})
	// Check error
	if err != nil {
		log.Println("error: ", err)
		Helper.Logger(err, false)
	} else {
		Rowefffect, _ := userData.RowsAffected()
		if Rowefffect == 0 {
			log.Println("input value is not change with previous one or id= " + fmt.Sprint(usr.Id) + "is not valid")
		}
		return Rowefffect > 0, err
	}
	return false, err
}
