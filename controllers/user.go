package controllers

import (
	"fmt"
	"net/http"
	"time"

	"assingmentTask.com/userManagement/configs"
	"assingmentTask.com/userManagement/models"
	"assingmentTask.com/userManagement/responses"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser(c echo.Context) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var user models.User
    defer cancel()

    //validate the request body
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
    }

    //use the validator library to validate required fields
    if validationErr := validate.Struct(&user); validationErr != nil {
        return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
    }

    newUser := models.User{
        Id:       primitive.NewObjectID(),
        Name:     user.Name,
    }

    result, err := userCollection.InsertOne(ctx, newUser)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
    }
    return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func GetAUser(c echo.Context) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    userId := c.Param("id")
    var user models.User
    defer cancel()

    fmt.Println("usersid",userId)
    objId, error := primitive.ObjectIDFromHex(userId)
    fmt.Println(objId)
    fmt.Println(error)
    err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
    }

    return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": user}})
}

func GetAllUsers(c echo.Context) error {
    fmt.Println("Inside get all users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var users []models.User
    defer cancel()

    results, err := userCollection.Find(ctx, bson.M{})

    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
    }

    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleUser models.User
        if err = results.Decode(&singleUser); err != nil {
            return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
        }

        fmt.Println("line 84",singleUser,singleUser.Id.Hex(),singleUser.Id)
        users = append(users, singleUser)
    }

    return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": users}})
}