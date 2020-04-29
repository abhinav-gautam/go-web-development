package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age int
}

func (p person) ShowAge() int {
	return p.Age
}
func (p person) DblAge() int {
	return p.Age * 2
}
func (p person) TakeArg(x int) int{
	return x * 2
}
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	p := person{"Abhinav", 22}

	err := tpl.Execute(os.Stdout,p)
	if err != nil {
		log.Fatalln(err)
	}
}
