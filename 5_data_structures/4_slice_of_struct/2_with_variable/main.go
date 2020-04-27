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
	Name string
	Motto string
}

func main() {
	g1 := graphics{"Nvidia","The way its meant to be played."}
	g2 := graphics{"EA Sports","To the game"}
	g3 := graphics{"AMD","The Future is Fusion"}

	gi := []graphics{g1,g2,g3}

	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",gi)
	if err != nil {
		log.Fatalln(err)
	}
}
