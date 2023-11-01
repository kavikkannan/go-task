package routes

import (
	"github.com/gorilla/mux"
	"github.com/kavikkannan/go-task/pkg/controllers"

)

var Registeruser = func (router *mux.Router)  {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/grievances/", controllers.CreateGrievances).Methods("POST")
	router.HandleFunc("/user/", controllers.GetUser).Methods("GET")
	router.HandleFunc("/employe/", controllers.GetEmploye).Methods("GET")
	router.HandleFunc("/grievances/", controllers.GetGrievances).Methods("GET")
	router.HandleFunc("/user/{useremail}", controllers.GetUserByEmail).Methods("GET")
	router.HandleFunc("/employe/{useremail}", controllers.GetEmployeByEmail).Methods("GET")
/* 	router.HandleFunc("/grievances/{useremail}", controllers.GetEmployeByGrievances).Methods("GET")
 */	router.HandleFunc("/user/{useremail}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/allgrievances/{useremail}", controllers.GetallGrievancesByEmail).Methods("GET")
	router.HandleFunc("/Ugrievances/{id}", controllers.UpdateGrievances).Methods("PUT")
	router.HandleFunc("/user/{useremail}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/grievances/{id}", controllers.DeleteGrievances).Methods("DELETE")
}