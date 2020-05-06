package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",set)
	http.HandleFunc("/read",read)
	http.HandleFunc("/abundance",abundance)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080",nil))

}
func set(w http.ResponseWriter,r *http.Request){
	http.SetCookie(w,&http.Cookie{
		Name:"id",
		Value:"49681321",
	})
	fmt.Fprint(w,"Cookies Written")

}
func read(w http.ResponseWriter,r *http.Request){
	c1,err := r.Cookie("id")
	if err != nil {
		http.Error(w,http.StatusText(400),400)
	}else{
		fmt.Fprint(w,"id: ",c1.Value,"\n")
	}
	c2,err := r.Cookie("general")
	if err != nil {
		http.Error(w,http.StatusText(400),400)
	}else{
		fmt.Fprint(w,"general: ",c2.Value,"\n")
	}
	c3,err := r.Cookie("specific")
	if err != nil {
		http.Error(w,http.StatusText(400),400)
	}else{
		fmt.Fprint(w,"specific: ",c3.Value,"\n")
	}
	c4,err := r.Cookie("user_id")
	if err != nil {
		http.Error(w,http.StatusText(400),400)
	}else{
		fmt.Fprint(w,"user_id: ",c4.Value,"\n")
	}

}
func abundance(w http.ResponseWriter,r *http.Request){
	http.SetCookie(w,&http.Cookie{
		Name:"general",
		Value:"abhinav",
	})
	http.SetCookie(w,&http.Cookie{
		Name:"specific",
		Value:"hacker",
	})
	http.SetCookie(w,&http.Cookie{
		Name:"user_id",
		Value:"1",
	})
	fmt.Fprint(w,"Cookies Written")
}
