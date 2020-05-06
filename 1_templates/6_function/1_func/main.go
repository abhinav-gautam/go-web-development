package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

var fm = template.FuncMap{
	"tu": strings.ToUpper,
	"ft": ft,
}

func ft(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

type graphics struct {
	Name  string
	Motto string
}
type cars struct {
	Brand string
	Name  string
}

func main() {
	g1 := graphics{"Nvidia", "The way its meant to be played."}
	g2 := graphics{"EA Sports", "To the game"}
	g3 := graphics{"AMD", "The Future is Fusion"}

	gi := []graphics{g1, g2, g3}

	c1 := cars{"Koenigsegg", "Agera"}
	c2 := cars{"Bugati", "Chiron"}
	c3 := cars{"Salen", "S7"}

	ci := []cars{c1, c2, c3}

	data := struct {
		Cars     []cars
		Graphics []graphics
	}{
		Cars:     ci,
		Graphics: gi,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
