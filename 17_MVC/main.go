package main

import (
	"abhinav-web-dev/17_MVC/controllers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id",uc.GetUser)
	r.POST("/user",uc.CreateUser)
	r.DELETE("/user/:id",uc.DeleteUser)
	log.Fatalln(http.ListenAndServe(":8080",r))
}

