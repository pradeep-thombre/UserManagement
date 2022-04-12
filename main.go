package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	// "github.com/labstack/echo"
	"gopkg.in/mgo.v2"
    "assingmentTask.com/userManagement/controllers"
)


func main() {

    router := httprouter.New()
    uc := controllers.NewUserController(getSession())
    // router.GET("/user", getUsers)
    router.GET("/user/:id", uc.GetUser)
    router.POST("/user", uc.CreateUser)
    http.ListenAndServe("localhost:3000",router)
}

func getSession() *mgo.Session{

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil{
		panic(err)
	}
	return s
}