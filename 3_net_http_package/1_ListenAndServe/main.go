package main

import (
	"fmt"
	"net/http"
)

type hotDog int

func (h hotDog) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Any code in this function")
}
func main() {
	var d hotDog
	http.ListenAndServe(":8080",d)
}
