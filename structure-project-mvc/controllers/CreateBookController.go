package controllers

import (
	"net/http"
	"projects/lib/database"

	"github.com/labstack/echo"
)

func CreateBookController(c echo.Context) error {
	book, err := database.CreateBook(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  book,
	})
}
