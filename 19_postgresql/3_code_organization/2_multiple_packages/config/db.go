package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB,err = sql.Open("postgres","postgres://abhinav:91189@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Database Connected")
}
