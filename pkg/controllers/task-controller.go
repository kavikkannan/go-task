package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/BalkanID-University/go-task/pkg/models"
	"github.com/BalkanID-University/go-task/pkg/utils"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUser()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetEmploye(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllEmploye()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userEmail := vars["useremail"]
    fmt.Println(userEmail)
    userDetails, _ := models.GetUserByEmail(userEmail)
    res, _ := json.Marshal(userDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetEmployeByEmail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userEmail := vars["useremail"]
    fmt.Println(userEmail)
    userDetails, _ := models.GetEmployeByEmail(userEmail)
    res, _ := json.Marshal(userDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userid"]
	fmt.Println(userId)
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsinf")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    var updateUser = &models.User{}
    utils.ParseBody(r, updateUser)
    vars := mux.Vars(r)
    userEmail := vars["useremail"] // Update the route variable name
    userDetails, db := models.GetUserByEmail(userEmail) // Update the function name

    if userDetails != nil {
        if updateUser.Project != "" {
            userDetails.Project = updateUser.Project
        }

        if updateUser.Task != "" {
            userDetails.Task = updateUser.Task
        }

        if updateUser.Email != "" {
            userDetails.Email = updateUser.Email
        }

        if updateUser.Deadline != "" {
            userDetails.Deadline = updateUser.Deadline
        }

        db.Save(&userDetails)
        res, _ := json.Marshal(userDetails)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(res)
    } else {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("User not found"))
    }
}

