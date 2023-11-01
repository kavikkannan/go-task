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
type Login struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Position string  `json:"position"`
	Password []byte `json:"-"`
}

type Grievances struct {
	Id       uint   `json:"id"`
	Email string `json:"email"`
	G_Type	string `json:"g_type"`
	Description string `json:"description"`
	Status    string `json:"status" `
	Date string  `json:"date"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Grievances{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func (c *Grievances) CreateGrievances() *Grievances {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetAllUser() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetAllGrievances() []Grievances {
	var Users []Grievances
	db.Find(&Users)
	return Users
}

func GetAllEmploye() []Login {
	var Users []Login
	db.Find(&Users)
	return Users
}

func GetAllEmployeByGrievances(userEmail string) ([]Grievances, error) {
    var userDetails []Grievances

    // Query the database using GORM to retrieve all rows associated with the given email
    if err := db.Where("email = ?", userEmail).Find(&userDetails).Error; err != nil {
        return nil, err
    }

    return userDetails, nil
}


func GetUserByEmail(email string) (*User, *gorm.DB) {
    var getUser User
    db := db.Where("email = ?", email).Find(&getUser)
    return &getUser, db
}

func GetEmployeByEmail(email string) (*Login, *gorm.DB) {
    var getUser Login
    db := db.Where("email = ?", email).Find(&getUser)
    return &getUser, db
}

func GetEmployeByGrievance(ID int64) (*Grievances, *gorm.DB) {
    var getUser Grievances
    db := db.Where("ID=?", ID).Find(&getUser)
    return &getUser, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}
func DeleteGrievances(ID int64) Grievances {
	var user Grievances
	db.Where("ID=?", ID).Delete(user)
	return user
}

