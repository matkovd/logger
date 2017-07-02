package dao

import (
	"database/sql"
	"../models"
	"fmt"
)

const SQL_CONNECT_STRING = "root:test@tcp(localhost)/logs"

func Insert(log models.Log) {
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

func GetByID(id int64) models.Log {
	db, err := sql.Open("mysql", SQL_CONNECT_STRING)
	if err != nil {
		panic(err)
	}
	query := fmt.Sprintf("Select From,Status,Message from logs where ID=%d);", id)
	rows, err := db.Query(query)
	var from string
	var status string
	var message string
	for rows.Next() {
		if err := rows.Scan(&from, &status, &message); err != nil {
			panic(err)
		}
	}
	res := models.Log{id, from, status, message}
	return res
}