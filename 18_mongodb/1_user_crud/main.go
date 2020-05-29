package __user_crud

import (
	"abhinav-web-dev/18_mongodb/controllers"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id",uc.GetUser)
	r.POST("/user",uc.CreateUser)
	r.DELETE("/user/:id",uc.DeleteUser)
	log.Fatalln(http.ListenAndServe(":8080",r))
}
func getSession() *mgo.Session{
	// Connect to our local mongo
	s,err := mgo.Dial("mongodb://localhost")

	// Check if connection error
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Database Connected")
	return s
}

