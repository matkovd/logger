package main

import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func data(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{'data':'working'}")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/data", data)
	http.ListenAndServe(":8080", nil)
}
