package main

import (
	c "github.com/iswanulumam/go-training-restful/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo{
	e := echo.New()

	//routing
	e.GET("/users", c.GetAllUsers)
	e.GET("/users/:id", c.GetUser)
	e.POST("/users", c.CreateUser)
	e.DELETE("/users/:id", c.DeleteUser)
	e.PUT("/users/:id", c.UpdateUser)

	return e
}

func main() {
	
}