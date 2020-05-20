package controllers

import (
	"abhinav-web-dev/17_MVC/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type UserController struct {}

func NewUserController() *UserController{
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	u := models.User{
		Name:"Abhinav Gautam",
		Gender:"Male",
		Age:25,
		Id:params.ByName("id"),
	}
	uj,err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"%s\n",uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter,r *http.Request,_ httprouter.Params){
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id="29"
	uj,err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type","application.json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w,"%s\n",uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	w.WriteHeader(200)
	fmt.Fprint(w,"Code to delete user")
}