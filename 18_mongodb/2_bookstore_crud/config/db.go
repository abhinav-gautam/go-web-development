package config

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

var DB *mgo.Database
var Books *mgo.Collection

func init() {
	var err error
	session,err := mgo.Dial("mongodb://abhinav:91189@localhost/bookstore")
	if err != nil {
		panic(err)
	}
	if err := session.Ping(); err!=nil{
		panic(err)
	}
	DB = session.DB("bookstore")
	Books = DB.C("books")
	fmt.Println("[+] Database Connected")
}
