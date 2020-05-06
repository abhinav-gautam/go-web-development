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

type graphics struct {
	Name  string
	Motto string
}

func main() {
	g1 := graphics{"Nvidia", "The way its meant to be played."}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", g1)
	if err != nil {
		log.Fatalln(err)
	}
}
