package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter,req *http.Request){
	fmt.Fprint(w,"Index")
}
func dog(w http.ResponseWriter,req *http.Request){
	fmt.Fprint(w,"Dog")
}
func me(w http.ResponseWriter,r *http.Request){
	tpl,err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(w,"Abhinav")
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	http.Handle("/",http.HandlerFunc(index))
	http.Handle("/dog/",http.HandlerFunc(dog))
	http.Handle("/me/",http.HandlerFunc(me))

	http.ListenAndServe(":8080",nil)
}

