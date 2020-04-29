package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	scores := struct {
		Score1 int
		Score2 int
	}{3,5}

	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",scores)
	if err != nil {
		log.Fatalln(err)
	}
}