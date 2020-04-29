package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fDMY": dayMonthYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func dayMonthYear(t time.Time) string{
	s := t.Format("02-01-2006")
	return s
}
func main() {
	t := time.Now()

	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",t)

	if err != nil {
		log.Fatalln(err)
	}
}
