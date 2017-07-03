package dao

import (
	"../models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const SQL_CONNECT_STRING = "root:test@tcp(localhost:3306)/logs"

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
	defer db.Close()
}

func GetByID(id int64) models.Log {
	db, err := sql.Open("mysql", SQL_CONNECT_STRING)
	if err != nil {
		panic(err)
	}
	query := fmt.Sprintf("Select Address,Status,Message from logs where ID=%d;", id)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	var from string
	var status string
	var message string
	for rows.Next() {
		if err := rows.Scan(&from, &status, &message); err != nil {
			panic(err)
		}
	}
	res := models.Log{id, from, status, message}
	defer rows.Close()
	defer db.Close()
	return res
}

func GetByStatus(status string) []models.Log {
	db, err := sql.Open("mysql", SQL_CONNECT_STRING)
	if err != nil {
		panic(err)
	}
	query := fmt.Sprintf("Select ID,Address,Status,Message from logs where Status='%s';", status)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	var logs []models.Log
	var from string
	var st string
	var id int64
	var message string
	for rows.Next() {
		if err := rows.Scan(&id, &from, &st, &message); err != nil {
			panic(err)
		}
		logs = append(logs, models.Log{id, from, st, message})
	}
	defer rows.Close()
	defer db.Close()
	return logs
}

func GetByFrom(from string) []models.Log {
	db, err := sql.Open("mysql", SQL_CONNECT_STRING)
	if err != nil {
		panic(err)
	}
	query := fmt.Sprintf("Select ID,Address,Status,Message from logs where Address='%s';", from)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	var logs []models.Log
	var fr string
	var status string
	var id int64
	var message string
	for rows.Next() {
		if err := rows.Scan(&id, &fr, &status, &message); err != nil {
			panic(err)
		}
		logs = append(logs, models.Log{id, fr, status, message})
	}
	defer rows.Close()
	defer db.Close()
	return logs
}

func GetByMessage(message string) []models.Log {
	db, err := sql.Open("mysql", SQL_CONNECT_STRING)
	if err != nil {
		panic(err)
	}
	query := fmt.Sprintf("Select ID,Address,Status,Message from logs where Message='%s';", message)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	var logs []models.Log
	var from string
	var status string
	var id int64
	var mes string
	for rows.Next() {
		if err := rows.Scan(&id, &from, &status, &mes); err != nil {
			panic(err)
		}
		logs = append(logs, models.Log{id, from, status, mes})
	}
	defer rows.Close()
	defer db.Close()
	return logs
}
