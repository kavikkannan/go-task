package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kavikkannan/go-task/pkg/models"
	"github.com/kavikkannan/go-task/pkg/utils"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUser()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetGrievances(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllGrievances()
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

    userDetails, _ := models.GetUserByEmail(userEmail)
    res, _ := json.Marshal(userDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetEmployeByEmail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userEmail := vars["useremail"]

    userDetails, _ := models.GetEmployeByEmail(userEmail)
    res, _ := json.Marshal(userDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

/* func GetEmployeByGrievances(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userEmail := vars["useremail"]

    userDetails, _ := models.GetEmployeByGrievance(userEmail)
    res, _ := json.Marshal(userDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
} */

func GetallGrievancesByEmail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userEmail := vars["useremail"]

 
    // Assume userDetails is a slice of structs or a list of data associated with the given email
    // Fetch all rows associated with the email
    allDetails, err := models.GetAllEmployeByGrievances(userEmail)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    res, err := json.Marshal(allDetails)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

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

func CreateGrievances(w http.ResponseWriter, r *http.Request) {
	CreateGrievances := &models.Grievances{}
	utils.ParseBody(r, CreateGrievances)
	c := CreateGrievances.CreateGrievances()
	res, _ := json.Marshal(c)
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

func DeleteGrievances(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	fmt.Println(userId)
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsinf")
	}
	user := models.DeleteGrievances(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateGrievances(w http.ResponseWriter, r *http.Request) {
    var updateUser = &models.Grievances{}
    utils.ParseBody(r, updateUser)
    vars := mux.Vars(r)
    userId := vars["id"] // Update the route variable name
    ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsinf")
	}
	userDetails, db := models.GetEmployeByGrievance(ID) // Update the function name

    if userDetails != nil {
        

        if updateUser.Status != "" {
            userDetails.Status = updateUser.Status
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

