package main

import (
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("v")
	io.WriteString(w, "You searched:"+v)
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
