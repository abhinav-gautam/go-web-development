package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var tpl *template.Template

type user struct {
	Username string
	Name string
	Password []byte
	role string
}

type session struct {
	Username string
	LastActivity time.Time
}

var dbUser = map[string]user{}
var dbSession = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}
func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/dashboard",dashboard)
	http.HandleFunc("/sign_up",signUp)
	http.HandleFunc("/login",login)
	http.HandleFunc("/admin",admin)
	http.HandleFunc("/logout",logout)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))

}
func index(w http.ResponseWriter,r *http.Request){
	tpl.ExecuteTemplate(w,"index.gohtml",nil)
}
func dashboard(w http.ResponseWriter,r *http.Request){
	if !alreadyLoggedIn(w,r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	u := getUser(r)
	tpl.ExecuteTemplate(w,"dashboard.gohtml",u)
}
func admin(w http.ResponseWriter,r *http.Request){
	if !alreadyLoggedIn(w,r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	c,_:=r.Cookie("session")
	username:=dbSession[c.Value].Username
	u := dbUser[username]
	if u.role != "admin"{
		http.Error(w,"Only for Admins",http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w,"admin.gohtml",u)
}