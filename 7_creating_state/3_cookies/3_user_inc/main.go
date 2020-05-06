package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))
}
func index(w http.ResponseWriter,r *http.Request){
	cookie,err := r.Cookie("user_count")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:"user_count",
			Value:"0",
		}
	}

	count,err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w,cookie)
	_,err = io.WriteString(w,cookie.Value)
	if err != nil {
		http.Error(w,http.StatusText(500),500)
	}
}