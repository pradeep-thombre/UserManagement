package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
    "github.com/pradeep-thombre/UserManagement/controllers"
)

type userResponse struct {
    Message string `json:"user"`
}

func main() {

    router := httprouter.New()
    uc := controllers.NewUserController(getSession())
    router.GET("/user", GetAllUsers)
    router.GET("/user/:id", uc.GetUser)
    router.POST("/user", uc.CreateUser)
    // router.DELETE("/user/:id", uc.DeleteUser)
    http.ListenAndServe("localhost:3000",router)
    // router.Logger.Fatal(r.Start(":3000"))
}

func GetAllUsers(c echo.Context) error {
    return c.JSON(http.StatusOK, userResponse{
        Message: "Pradeep",
    })
}

func getSession() *mgo.Session{
    s , err :=mgo.Dial("mongodb://localhost:27017")
    if err != nil{
        panic(err)
    }
    return s
}