package main

import (
	"fmt"
	"log"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w,"Dog dog dog")
}
func c(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w,"Cat cat cat")
}

func main() {
	// Here http.HandleFunc is a function that takes (string, func(http.ResponseWriter,*http.Request))
	http.HandleFunc("/dog/",d)
	http.HandleFunc("/cat",c)

	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatalln(err)
	}
}
