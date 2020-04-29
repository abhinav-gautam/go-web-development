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
type users struct {
	Name string
	Admin bool
}
func main() {
	u1 := users{"Abhinav",true}
	u2 := users{"Ansh",false}
	u3 := users{"",true}

	us := []users{u1,u2,u3}

	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",us)
	if err != nil {
		log.Fatalln(err)
	}


}
