package models

import (
	"log"
	"reakgo/utility"

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
	query := "SELECT orders.id,orders.productId,orders.placedOn,orders.expiry,orders.price,orders.buyer,orders.status FROM `orders` INNER JOIN authentication on authentication.id=orders.buyer WHERE 1=1 " + filter.WhereCondition + " ORDER BY orders.id DESC;"
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
		params.WhereCondition += " AND authentication.companyId= :companyId"
	}
	return params
}

// type Orders struct {
// 	Id        int64  `json:"id" db:"id" `
// 	ProductId int64  `json:"productId" db:"productId" `
// 	PlacedOn  int64  `json:"placedOn" db:"placedOn" `
// 	Expiry    int64  `json:"expiry" db:"expiry" `
// 	Price     int64  `json:"price" db:"price" `
// 	Buyer     int64  `json:"buyer" db:"buyer" `
// 	Status    string `json:"status" db:"status" `
// }

func (Orders) PostOrder(order Orders, tx *sqlx.Tx) (bool, error) {
	row, err := tx.NamedExec("UPDATE `orders` SET productId=:ProductId,placedOn=:PlacedOn,expiry=:Expiry,price=:Price,buyer=:Buyer,status=:Status WHERE id=:Id ", map[string]interface{}{"ProductId": order.ProductId, "PlacedOn": order.PlacedOn, "Expiry": order.Expiry, "Price": order.Price, "Buyer": order.Buyer, "Status": order.Status, "Id": order.Id})
	if err != nil {
		log.Println(err)
		return false, err
	}
	rowUpdate, _ := row.RowsAffected()
	return rowUpdate > 0, nil
}

func (Orders) GetSingleProduct(productId int64, tx *sqlx.Tx) (int64, error) {
	var planValidity int64
	err := tx.Get(&planValidity, "SELECT `plan_validity` FROM `products` WHERE  id=?", productId)
	return planValidity, err
}

func (Orders) PutOrder(data Orders, tx *sqlx.Tx) (bool, error) {
	_, err := tx.NamedExec("INSERT INTO `orders` (productId,placedOn,expiry,price,buyer,status) VALUES (:ProductId,:PlacedOn,:Expiry,:Price,:Buyer,:Status)", map[string]interface{}{"ProductId": data.ProductId, "PlacedOn": data.PlacedOn, "Expiry": data.Expiry, "Price": data.Price, "Buyer": data.Buyer, "Status": data.Status})
	if err != nil {
		log.Println(err)
		return true, err
	}
	return false, nil
}

func (Orders) OrderDelete(id int) (bool, error) {
	row, err := utility.Db.Exec("DELETE FROM orders WHERE id = ?", id)
	if err != nil {
		log.Print(err)
		return false, err
	}
	rowsDeleted, err := row.RowsAffected()
	if err != nil {
		log.Print(err)
		return false, err
	}
	return rowsDeleted > 0, nil
}
