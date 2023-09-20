package models

import (
	"github.com/jmoiron/sqlx"
)

type Orders struct {
	Id        int64  `db:"id" json:"id"`
	ProductId int64  `db:"productId" json:"productId"`
	PlacedOn  int64  `db:"placedOn" json:"placedOn"`
	Expiry    int64  `db:"expiry" json:"expiry"`
	Price     int64  `db:"price" json:"price"`
	Buyer     int64  `db:"buyer" json:"buyer"`
	Status    string `db:"status"`
}

type OrdersModel struct {
	DB *sqlx.DB
}

func (ord OrdersModel) GetOrders(data Orders) (bool, error) {

	return false, nil
}
