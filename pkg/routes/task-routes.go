package routes

import (
	"github.com/gorilla/mux"
	"github.com/kavikkannan/go-task/pkg/controllers"

)

var Registeruser = func (router *mux.Router)  {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userid}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userid}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userid}", controllers.DeleteUser).Methods("DELETE")
}