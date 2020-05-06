package main

import (
	uuid "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type user struct {
	Username string
	FName string
	LName string
}

var dbUser = map[string]user{}
var dbSession = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/form_process",formProcess)
	http.HandleFunc("/display",display)
	http.Handle("/facicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))
}
func index(w http.ResponseWriter,r *http.Request){
	cookie,err := r.Cookie("session")
	if err == http.ErrNoCookie {
		id := uuid.Must(uuid.NewV4())
		cookie = &http.Cookie{
			Name:"session",
			Value:id.String(),
		}
	}
	http.SetCookie(w,cookie)
	var u user
	if username,ok := dbSession[cookie.Value]; ok{
		u = dbUser[username]
	}
	tpl.ExecuteTemplate(w,"index.gohtml",u)
}
func formProcess(w http.ResponseWriter,r *http.Request){
	c,err := r.Cookie("session")
	if err == http.ErrNoCookie {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	r.ParseForm()
	username := r.FormValue("username")
	fName := r.FormValue("fName")
	lName := r.FormValue("lName")
	u := user{username,fName,lName}
	dbUser[username] = u
	dbSession[c.Value] = username
	http.Redirect(w,r,"/",http.StatusSeeOther)
}
func display(w http.ResponseWriter,r *http.Request){
	c,err := r.Cookie("session")
	if err == http.ErrNoCookie {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	var u user
	username,ok := dbSession[c.Value]
	if !ok{
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	u = dbUser[username]

	tpl.ExecuteTemplate(w,"display.gohtml",u)
}
