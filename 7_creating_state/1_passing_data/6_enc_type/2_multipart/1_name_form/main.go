package __name_form

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func index(w http.ResponseWriter,r *http.Request){
	bs := make([]byte,r.ContentLength)
	r.Body.Read(bs)
	body:=string(bs)

	err := tpl.ExecuteTemplate(w,"index.gohtml",body)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}
func main() {
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))
}
