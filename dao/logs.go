package dao

import (
	"database/sql"
	"../models"
	"fmt"
)

const SQL_CONNECT_STRING = "root:test@tcp(localhost)/logs"

func insert(log models.Log) {
	db, err := sql.Open("mysql", SQL_CONNECT_STRING)
	if err != nil {
		panic(err)
	}
	query := fmt.Sprintf("Insert into logs values( %d, %s, %s, %s);", log.ID, log.From, log.Status, log.Message)
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	db.Close();
}
