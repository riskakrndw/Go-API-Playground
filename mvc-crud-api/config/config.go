package config

import (
	"projects/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {
	// connectionString := os.Getenv("CONNECTION_STRING")
	connectionString := "root:welcome12345@tcp(localhost:3306)/new?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

// func InitPort() {
// 	var err error
// 	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
// 	if err != nil {
// 		panic(err)
// 	}
// }

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
