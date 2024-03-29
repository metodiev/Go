package main 

import (
	"log"
	"net/http"
	"github/com/forilla/mux"
	_ "github.com/jinzhu/gorm/dialect/mysql"
	"github.com/akhil/go-bookstore/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}