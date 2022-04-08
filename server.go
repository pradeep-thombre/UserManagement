package main

import (
    "net/http"

    "github.com/labstack/echo"
)

type userResponse struct {
    Message string `json:"user"`
}

func main() {
    e := echo.New()
    e.GET("/user", getUsers)
    e.Logger.Fatal(e.Start(":3000"))
}

func getUsers(c echo.Context) error {
    return c.JSON(http.StatusOK, userResponse{
        Message: "Pradeep",
    })
}