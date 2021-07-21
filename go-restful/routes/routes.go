package routes

import (
	"alta/training/constants"
	"alta/training/controllers"
	"alta/training/middlewares"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	e.GET("/users", controllers.GetUserControllers)
	e.GET("/users/:id", controllers.GetUserDetailControllers)
	e.POST("/login", controllers.LoginUsersControllers)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users/:id", controllers.GetUserDetailControllers)

	// eauthGroup
	eAuth := e.Group("")
	eAuth.Use(middleware.BasicAuth(middlewares.BasicAuthDb))
	eAuth.GET("/rahasia", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "rahasia, cuma bisa diakses yang login",
		})
	})
}
