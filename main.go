package main

import (
	"./dao"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func data(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{'data':'working'}")
}

func findByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[1:])
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
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/id", findByID)
	http.ListenAndServe(":8080", nil)
}
