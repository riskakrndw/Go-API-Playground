package routes

import (
	"projects/constants"
	"projects/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	//routing
	//---users
	// e.GET("/users", controllers.GetUsersController)
	// e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	// e.DELETE("/users/:id", controllers.DeleteUserController)
	// e.PUT("/users/:id", controllers.UpdateUserController)

	e.POST("/login", controllers.LoginUsersController)

	//---jwt users
	ajwt := e.Group("")
	ajwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	ajwt.GET("/users", controllers.GetUsersController)
	ajwt.GET("/users/:id", controllers.GetUserController)
	ajwt.DELETE("/users/:id", controllers.DeleteUserController)
	ajwt.PUT("/users/:id", controllers.UpdateUserController)

	//---books
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
}
