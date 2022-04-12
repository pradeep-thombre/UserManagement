package main

import (
	"fmt"

	"assingmentTask.com/userManagement/configs"
	"assingmentTask.com/userManagement/routes"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	// "net/http"
)

func main() {

	router := echo.New()
	
	//run database
    configs.ConnectDB()

	//routes
    routes.UserRoute(router) 
	fmt.Println("Inside main")
	router.Start("localhost:3000")
}

