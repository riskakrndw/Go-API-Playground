package controllers

import (
	"net/http"
	"projects/lib/database"

	"github.com/labstack/echo"
)

func CreateUserController(c echo.Context) error {
	user, err := database.CreateUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}
