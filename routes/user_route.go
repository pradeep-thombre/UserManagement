package routes

import (
	"fmt"

	"assingmentTask.com/userManagement/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoute(router *echo.Echo)  {
    //All routes related to users comes here
	fmt.Println("inside user route")
	router.GET("/", func(c echo.Context) error {
		fmt.Println("Simple route")
		return nil
	})
	router.POST("/users", controllers.CreateUser) 
	router.GET("/users/:id", controllers.GetAUser)
	router.GET("/users", controllers.GetAllUsers)
}