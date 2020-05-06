package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type user struct {
	Username string
	Name string
	Password []byte
}

var dbUser = map[string]user{}
var dbSession = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/dashboard",dashboard)
	http.HandleFunc("/form_process",formProcess)
	http.HandleFunc("/logout",logout)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))

}
func index(w http.ResponseWriter,r *http.Request){
	if alreadyLoggedIn(r){
		http.Redirect(w,r,"/dashboard",http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w,"index.gohtml",nil)
}
func dashboard(w http.ResponseWriter,r *http.Request){
	if !alreadyLoggedIn(r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	u := getUser(r)
	tpl.ExecuteTemplate(w,"dashboard.gohtml",u)
}
func logout(w http.ResponseWriter,r *http.Request){
	c, err:=r.Cookie("session")
	if err == http.ErrNoCookie {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	c.MaxAge = -1
	http.SetCookie(w,c)
	http.Redirect(w,r,"/",http.StatusSeeOther)
}