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
type cars struct {
	Brand string
	Name string
}
type items struct {
	Cars []cars
	Graphics []graphics
}

func main() {
	g1 := graphics{"Nvidia","The way its meant to be played."}
	g2 := graphics{"EA Sports","To the game"}
	g3 := graphics{"AMD","The Future is Fusion"}

	gi := []graphics{g1,g2,g3}

	c1 := cars{"Koenigsegg","Agera"}
	c2 := cars{"Bugati","Chiron"}
	c3 := cars{"Salen","S7"}

	ci := []cars{c1,c2,c3}

	data := items{ci,gi}

	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",data)
	if err != nil {
		log.Fatalln(err)
	}
}
