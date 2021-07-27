package controller

import (
	"tdd/rest/model"

	"github.com/labstack/echo"
)

func CreateGetBookController(bookModel model.BookModel) echo.HandlerFunc {
	return func(c echo.Context) error {
		books := bookModel.Get()
		return c.JSON(200, books)
	}
}
