package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/dog",dog)
	http.HandleFunc("/cat",cat)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))
}

func cat(w http.ResponseWriter,r *http.Request){
	fmt.Println("Request Method at Cat: ",r.Method)
}

func dog(w http.ResponseWriter,r *http.Request){
	fmt.Println("Request Method at Dog: ",r.Method)
	http.Redirect(w,r,"/cat",http.StatusMovedPermanently)
}
