package main

import (
	"io"
	"log"
	"net/http"
)

func index (w http.ResponseWriter,r *http.Request){
	name := r.FormValue("name")
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,"Your name is: "+name)
	io.WriteString(w,`
	<form action="/" method="post">
		<input type="text" name="name" placeholder="Enter Name">
		<input type="submit" value="submit">
	</form>`)

}
func main() {
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))
}
