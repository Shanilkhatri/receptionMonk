package models

import (
	"log"
	"reakgo/utility"

	"github.com/jmoiron/sqlx"
)

type Products struct {
	Id           int64  `db:"id"`
	Name         string `db:"name"`
	Price        int64  `db:"price"`
	PlanValidity int64  `db:"plan_validity" json:"plan_validity"`
	Description  string `db:"description"`
	Status       string `db:"status"`
}

type ProductModel struct {
	DB *sqlx.DB
}

func (prod ProductModel) PutProducts(data Products) (bool, error) {
	_, err := utility.Db.NamedExec("INSERT INTO `products`(name,price,plan_validity,description,status) VALUES (:Name,:Price,:PlanValidity,:Description,:Status)", map[string]interface{}{"Name": data.Name, "Price": data.Price, "PlanValidity": data.PlanValidity, "Description": data.Description, "Status": data.Status})
	if err != nil {
		log.Println(err)
		return true, err

	}
	return false, nil
}
