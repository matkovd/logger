package models

type Log struct {
	ID      int64  `json: "id"`
	From    string `json: "from"`
	Status  string `json: "status"`
	Message string ` json: "message"`
}
