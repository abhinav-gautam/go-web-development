package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/process",process)
	http.HandleFunc("/success",success)
	http.HandleFunc("/failed",failed)
	http.Handle("/favicon",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))
}

func index(w http.ResponseWriter,r *http.Request){
	fmt.Println("Request Method at Index: ",r.Method)
	err := tpl.ExecuteTemplate(w,"index.gohtml",nil)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}

func process(w http.ResponseWriter,r *http.Request){
	fmt.Println("Request Method at Process: ",r.Method)
	// Form Processing
	r.ParseForm()
	name := r.FormValue("name")
	password := r.FormValue("password")
	if name == "Abhinav" && password == "123" {
		http.Redirect(w,r,"/success",http.StatusSeeOther)
	}else{
		http.Redirect(w,r,"/failed",http.StatusSeeOther)
	}

}

func success(w http.ResponseWriter,r *http.Request){
	fmt.Println("Request Method at Success: ",r.Method)
	err := tpl.ExecuteTemplate(w,"form_success.gohtml",nil)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}

func failed(w http.ResponseWriter,r *http.Request){
	fmt.Println("Request Method at Failed: ",r.Method)
	err := tpl.ExecuteTemplate(w,"form_failed.gohtml",nil)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}