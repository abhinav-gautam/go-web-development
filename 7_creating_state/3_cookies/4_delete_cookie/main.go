package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/set",set)
	http.HandleFunc("/read",read)
	http.HandleFunc("/expire",expire)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))


}
func index(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html charset=utf-8")
	fmt.Fprint(w,`<h4><a href="/set">Set Cookie</a></h4>`)
}
func set(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html charset=utf-8")
	http.SetCookie(w,&http.Cookie{
		Name:"SomeCookie",
		Value:"EDITH",
	})
	fmt.Fprint(w,`<br><h4><a href="/read">Read Cookie</a></h4>`)

}
func read(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html charset=utf-8")
	cookie, err := r.Cookie("SomeCookie")
	if err != nil {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	fmt.Fprint(w,cookie)
	fmt.Fprint(w,`<h4><a href="/expire">Delete Cookie</a></h4>`)
}
func expire(w http.ResponseWriter,r *http.Request){
	cookie, err := r.Cookie("SomeCookie")
	if err != nil {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(w,cookie)
	http.Redirect(w,r,"/",http.StatusSeeOther)
}