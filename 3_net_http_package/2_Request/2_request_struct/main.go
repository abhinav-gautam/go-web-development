package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method  string
		URL 	*url.URL
		Submissions map[string][]string
		Header http.Header
		Host string
		ContentLength int64

	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	err = tpl.ExecuteTemplate(w,"index.gohtml",data)
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	var d hotdog
	err := http.ListenAndServe(":8080",d)
	if err != nil {
		log.Fatalln(err)
	}
}
