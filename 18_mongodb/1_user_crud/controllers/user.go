package controllers

import (
	"abhinav-web-dev/18_mongodb/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	// Get id
	id := params.ByName("id")

	//Verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	// Get ObjectId from hex
	oid:=bson.ObjectIdHex(id)

	u:= models.User{}

	// Fetch user
	err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u)
	if err != nil {
		w.WriteHeader(404)
		return
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

	// Create bson id
	u.Id=bson.NewObjectId()

	// Store the user in MongoDB
	uc.session.DB("go-web-dev-db").C("users").Insert(u)
	fmt.Println("[+] User created")

	uj,err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type","application.json")
	w.WriteHeader(http.StatusCreated) //201
	fmt.Fprintf(w,"%s\n",uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	id := params.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid:=bson.ObjectIdHex(id)

	err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w,"[+] User Deleted")
}