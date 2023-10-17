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
	Status    string `db:"status" json:"status"`
	DateFrom  int64  `json:"date_from"`
	DateTo    int64  `json:"date_to"`
	UserId    int64  `json:"userId"`
	CompanyId int64  `json:"companyId"`
}
type GetOrderFilter struct {
	WhereCon  string
	Id        int
	CompanyId int
	UserId    int
	Date      int64
}
type OrderDataCondition struct {
	WhereCondition string
	Orders
}

func (ord Orders) GetOrders(filter OrderDataCondition, tx *sqlx.Tx) ([]Orders, error) {
	var ordersArr []Orders
	query := "SELECT orders.id,orders.productId,orders.placedOn,orders.expiry,orders.price,orders.buyer,orders.status FROM `orders` INNER JOIN users on users.id=orders.buyer WHERE 1=1 " + filter.WhereCondition + " ORDER BY orders.id DESC;"
	condtion := map[string]interface{}{
		"id":        filter.Id,
		"date_from": filter.DateFrom,
		"date_to":   filter.DateTo,
		"companyId": filter.CompanyId,
		"userId":    filter.UserId,
	}
	rows, err := tx.NamedQuery(query, condtion)
	if err != nil {
		return ordersArr, err
	}
	defer rows.Close()
	for rows.Next() {
		var singleRow Orders
		err := rows.Scan(&singleRow.Id, &singleRow.ProductId, &singleRow.PlacedOn, &singleRow.Expiry, &singleRow.Price, &singleRow.Buyer, &singleRow.Status)
		if err != nil {
			return ordersArr, err
		}

		ordersArr = append(ordersArr, singleRow)
	}
	return ordersArr, nil
}

// Building conditions for getting orders Data
func (ord Orders) GetParamsForFilterOrderData(params OrderDataCondition) OrderDataCondition {

	if params.Id != 0 {
		params.WhereCondition += " AND orders.id = :id"
	}
	if params.UserId != 0 {
		params.WhereCondition += " AND orders.buyer = :userId"
	}
	if params.DateFrom != 0 && params.DateTo != 0 {
		params.WhereCondition += " AND orders.placedOn BETWEEN :date_from AND :date_to"
	} else if params.DateFrom != 0 {
		params.WhereCondition += " AND orders.placedOn >= :date_from"
	} else if params.DateTo != 0 {
		params.WhereCondition += " AND orders.placedOn <= :date_to"
	} else if params.CompanyId != 0 {
		params.WhereCondition += " AND users.companyId= :companyId"
	}
	return params
}
