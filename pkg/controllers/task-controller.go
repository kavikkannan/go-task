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

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userid"]
	fmt.Println(userId)
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsinf")
	}
	userDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
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
	userId := vars["bookid"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, db := models.GetUserById(ID)
	if updateUser.Name != "" {
		userDetails.Name = updateUser.Name
	}

	if updateUser.Publication != "" {
		userDetails.Publication = updateUser.Publication
	}

	if updateUser.Author != "" {
		userDetails.Author = updateUser.Author
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
