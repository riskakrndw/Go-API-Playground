package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name string
	Age  int
}

func GetUserController(c echo.Context) error {
	// return nil
	users := []User{
		{Name: "Ris", Age: 1},
		{Name: "Ka", Age: 1},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"users":   users,
	})
}
