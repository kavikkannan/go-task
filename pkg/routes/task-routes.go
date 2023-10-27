package routes

import (
	"github.com/gorilla/mux"
	"github.com/BalkanID-University/go-task/pkg/controllers"

)

var Registeruser = func (router *mux.Router)  {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetUser).Methods("GET")
	router.HandleFunc("/employe/", controllers.GetEmploye).Methods("GET")
	router.HandleFunc("/user/{useremail}", controllers.GetUserByEmail).Methods("GET")
	router.HandleFunc("/employe/{useremail}", controllers.GetEmployeByEmail).Methods("GET")
	router.HandleFunc("/user/{useremail}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{useremail}", controllers.DeleteUser).Methods("DELETE")
}