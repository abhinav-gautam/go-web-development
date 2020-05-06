package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))
}

func index(w http.ResponseWriter,r *http.Request){
	cookie,err := r.Cookie("session")

	if err == http.ErrNoCookie {
		id := uuid.Must(uuid.NewV4())
		cookie = &http.Cookie{
			Name:"session",
			Value:id.String(),
			//Secure:true,
			HttpOnly:true,
		}
		http.SetCookie(w,cookie)
	}

	fmt.Println(cookie)
}