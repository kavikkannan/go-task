package models

import (
	"github.com/jinzhu/gorm"
	"github.com/kavikkannan/go-task/pkg/config"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Project        string `gorm:"" json:"project"`
	Email      string `json:"email"`
	Task	 string `json:"task"`
	Deadline	 string `json:"deadline"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllUser() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserByEmail(email string) (*User, *gorm.DB) {
    var getUser User
    db := db.Where("email = ?", email).Find(&getUser)
    return &getUser, db
}
func DeleteUser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}
