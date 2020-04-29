package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter,req *http.Request){
	fmt.Fprint(w,"Index")
}
func dog(w http.ResponseWriter,req *http.Request){
	fmt.Fprint(w,"Dog")
}
func me(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"I am Abhinav Gautam")
}
func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/dog/",dog)
	http.HandleFunc("/me/",me)

	http.ListenAndServe(":8080",nil)
}
