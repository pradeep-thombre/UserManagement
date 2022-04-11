package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"assingmentTask.com/userManagement/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/x/mongo/driver/session"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    "github.com/pradeep-thombre/UserManagement/models"
)
type Usercontroller struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *Usercontroller{
	return &Usercontroller{s}
}

func (uc Usercontroller) GetUser(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	id := param.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}
	err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err!=nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n",uj)
}

func (uc Usercontroller) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id= bson.NewObjectId()
	uc.session.DB("mongo-golang").C('users').Insert(u)
	uj,err := json.Marshal(u)
	if err !=nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w,"%s\n",uj)
}