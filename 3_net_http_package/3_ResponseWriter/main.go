package main

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog int
func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Abhinav-Key","This is my key")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w,"<h1>Go suffer no fools</h1>")
}
func main() {
	var d hotdog
	err := http.ListenAndServe(":8080",d)
	if err != nil {
		log.Fatalln(err)
	}
}
