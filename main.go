package main

import (
	"./dao"
	"./models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working")
}

func insert(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var log models.Log
	err := decoder.Decode(&log)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	dao.Insert(log)
	w.Write([]byte("OK"))
}

func findByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[4:])
	fmt.Fprintf(w, string(id))
	if err != nil {
		panic(err)
	}
	logs := dao.GetByID(int64(id))
	logsMarshalled, err := json.Marshal(logs)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(logsMarshalled))
}

func findByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Path[8:]
	logs := dao.GetByStatus(status)
	logsMarshalled, err := json.Marshal(logs)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(logsMarshalled))
}

func findByMessage(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path[9:]
	logs := dao.GetByMessage(message)
	logsMarshalled, err := json.Marshal(logs)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(logsMarshalled))
}

func findByFrom(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Path[6:]
	logs := dao.GetByFrom(from)
	logsMarshalled, err := json.Marshal(logs)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(logsMarshalled))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/id/", findByID)
	http.HandleFunc("/from/", findByFrom)
	http.HandleFunc("/status/", findByStatus)
	http.HandleFunc("/insert/", insert)
	http.HandleFunc("/message/", findByMessage)
	http.ListenAndServe(":8080", nil)
}
