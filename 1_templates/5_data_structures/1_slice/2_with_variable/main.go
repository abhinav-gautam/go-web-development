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
	friends := []string{"Ansh", "Jayant", "Manasi", "Nimo"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", friends)
	if err != nil {
		log.Fatalln(err)
	}
}
