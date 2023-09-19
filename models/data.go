package models

import (
	"log"
	"reakgo/utility"

	"github.com/jmoiron/sqlx"
)

type Data struct {
	Id   int32
	Name string
}

type DataModel struct {
	DB *sqlx.DB
}

func (data DataModel) All() ([]Data, error) {
	var allRows []Data

	rows, err := utility.Db.Query("SELECT * FROM data")
	if err != nil {
		log.Println("Query failed")
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var singleRow Data
		err := rows.Scan(&singleRow.Id, &singleRow.Name)
		if err != nil {
			log.Println("Query Scan failed")
			log.Println(err)
		}
		allRows = append(allRows, singleRow)
	}
	if err = rows.Err(); err != nil {
		log.Println("Couldn't fetch from data")
		log.Println(err)
	}

	return allRows, err
}
