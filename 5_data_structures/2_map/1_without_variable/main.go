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
	cars := map[string]string{
		"Koenigsegg":"Agera",
		"Bugati":"Chiron",
		"Lamborgini":"Elemento",
		"Salen":"S7",
	}

	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",cars)
	if err != nil {
		log.Fatalln(err)
	}
}
