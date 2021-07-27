package main

import (
	"tdd/rest/config"
	"tdd/rest/controller"
	"tdd/rest/model"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	db, err := gorm.Open(mysql.Open(config.DB_CONNECTION_STRING))
	db.AutoMigrate(&model.Book{})
	if err != nil {
		panic(err)
	}
	bookModel := model.NewGormBookModel(db)

	//get
	bookController := controller.CreateGetBookController(bookModel)
	e.GET("/book", bookController)

	//post
	postBookController := controller.CreatePostBookController(bookModel)
	e.POST("/book", postBookController)

	e.Start(":8080")
}
