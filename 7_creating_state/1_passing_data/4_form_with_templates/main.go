package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func index(w http.ResponseWriter,r *http.Request){
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	data := struct {
		Fname string
		Lname string
	}{fname,lname}
	tpl.ExecuteTemplate(w,"index.gohtml",data)
}
func main() {
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))
}