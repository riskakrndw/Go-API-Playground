package controllers

import (
	"net/http"
	"projects/lib/database"
	"projects/models"
	"strconv"

	"github.com/labstack/echo"
)

func UpdateBookController(c echo.Context) error {
	var book models.Book
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	c.Bind(&book)
	update_book, err := database.UpdateBook(id, book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    update_book,
	})
}
