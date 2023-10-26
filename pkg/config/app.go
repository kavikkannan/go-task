package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db* gorm.DB
)

func Connect(){
	d, err:= gorm.Open("mysql","admin:subi1234@tcp(database-1.c5akcsdv8dty.ap-south-1.rds.amazonaws.com)/task?charset=utf8&parseTime=True&loc=Local")	
	if err != nil{
		panic(err)
	}
	db=d
}

func GetDB() *gorm.DB{
	return db
}