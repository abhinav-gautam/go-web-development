package main

import (
	"fmt"
	"log"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w,"dog dog dog")
}

func c(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w,"cat cat cat")
}

func main() {
	// Here http.HandlerFunc is a type of func(http.ResponseWriter,*http.Request), function ServeHTTP(w ResponseWriter, r *Request)
	// is attached to it therefore it also implements Handler interface. Type casting is happening here.

	http.Handle("/dog/",http.HandlerFunc(d))
	http.Handle("/cat",http.HandlerFunc(c))

	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatalln(err)
	}
}
