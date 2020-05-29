package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)
var db *sql.DB

func init() {
	var err error
	db,err = sql.Open("postgres","postgres://abhinav:91189@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Database Connected")
}
