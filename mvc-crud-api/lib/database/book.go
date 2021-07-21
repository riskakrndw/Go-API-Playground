package database

import (
	"projects/config"
	"projects/models"

	"github.com/labstack/echo"
)

func CreateBook(c echo.Context) (interface{}, error) {
	book := models.Book{}
	c.Bind(&book)
	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func GetBooks() (interface{}, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBook(id int) (interface{}, error) {
	var book models.Book
	if err := config.DB.Find(&book, "id=?", id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(id int, book interface{}) (interface{}, error) {
	if err := config.DB.Find(&book, "id=?", id).Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id int) (interface{}, error) {
	var book models.Book
	if err := config.DB.Find(&book, "id=?", id).Delete(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}
