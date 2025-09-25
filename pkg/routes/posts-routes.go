package routes

import (
	"github.com/gorilla/mux"
	"github.com/guivari/sharingvision-backend/pkg/controllers"

)

var RegisterPostsRoutes = func(router *mux.Router) {
	router.HandleFunc("/article/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/article/{limit}/{offset}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/article/{id}", controllers.GetPostByID).Methods("GET")
	router.HandleFunc("/article/{id}", controllers.DeletePost).Methods("DELETE")
	router.HandleFunc("/article/{id}", controllers.UpdatePost).Methods("PUT")

}