package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServeTLS(":10443","cert.pem","key.pem",nil))
}
func index(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"This is TLS Server")
}
