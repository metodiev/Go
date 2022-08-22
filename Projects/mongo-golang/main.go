package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/metodiev/Go/Projects/mongo/golang/controllers"
	"github.com/metodiev/Go/Projects/mongo/golang/models"

)

func main()  {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() * mgo.getSession{
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return s
}