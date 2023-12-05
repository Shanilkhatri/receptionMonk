package models

import "reakgo/utility"

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
