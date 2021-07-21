package controllers

import (
	"net/http"
	"projects/lib/database"
	"projects/models"
	"strconv"

	"github.com/labstack/echo"
)

func UpdateUserController(c echo.Context) error {
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	c.Bind(&user)
	update_user, err := database.UpdateUser(id, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    update_user,
	})
}
