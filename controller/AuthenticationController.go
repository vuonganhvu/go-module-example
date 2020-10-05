package controller

import (
	"database/sql"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	r.Cookies()
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/go-example")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

}
