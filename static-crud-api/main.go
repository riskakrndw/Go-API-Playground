package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users = map[int]*User{}
var no = 1

// --------------- controller -----------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// //get user by id
func GetUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    users[userId],
	})
}

// //delete user by id
func DeleteUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))
	delete(users, userId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success deleted user",
	})
}

// //update user by id
func UpdateUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))

	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	users[userId].Name = user.Name
	users[userId].Email = user.Email
	users[userId].Password = user.Password
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

// func UpdateUserController(c echo.Context) error {
// 	user := new(User)
// 	if err := c.Bind(user); err != nil {
// 		return err
// 	}
// 	userId, _ := strconv.Atoi(c.Param("userId"))
// 	users[userId].Name = user.Name
// 	users[userId].Email = user.Email
// 	users[userId].Password = user.Password
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success update user",
// 		"user":    users[userId],
// 	})
// }

//create new user
func CreateUserController(c echo.Context) error {
	//binding data
	u := &User{
		Id: no,
	}
	c.Bind(&u)

	users[u.Id] = u
	no++
	return c.JSON(http.StatusOK, u)
}

// ----------------------------------------------

func main() {
	e := echo.New()

	//routing with query param
	e.GET("/users", GetUsersController)
	e.GET("/user/:userId", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:userId", UpdateUserController)
	e.DELETE("/users/:userId", DeleteUserController)

	//start the server and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
