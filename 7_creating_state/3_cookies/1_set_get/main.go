package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/set",set)
	http.HandleFunc("/read",read)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))
}

func set(w http.ResponseWriter,r *http.Request){
	http.SetCookie(w,&http.Cookie{
		Name:"CookieName",
		Value:"Cookie Value",
	})
	_,err := io.WriteString(w,"Cookies Written")
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	}

}
func read(w http.ResponseWriter,r *http.Request){
	c,err := r.Cookie("CookieName")
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	}
	_,err = fmt.Fprint(w,"Your Cookie: ",c)
	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	}
}