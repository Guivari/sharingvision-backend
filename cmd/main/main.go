package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/guivari/sharingvision-backend/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterPostsRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}