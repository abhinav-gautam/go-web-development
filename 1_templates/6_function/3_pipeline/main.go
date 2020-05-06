package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fd":    double,
	"fsqr":  square,
	"fsqrt": sqrt,
}

func double(x float64) float64 {
	return x * 2
}
func square(x float64) float64 {
	return x * x
}
func sqrt(x float64) float64 {
	return math.Sqrt(x)
}
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", float64(6))
	if err != nil {
		log.Fatalln(err)
	}
}
