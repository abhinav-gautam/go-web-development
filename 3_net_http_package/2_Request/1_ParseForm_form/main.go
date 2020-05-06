package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	var d hotdog
	err := http.ListenAndServe(":8080", d)
	if err != nil {
		log.Fatalln(err)
	}
}
