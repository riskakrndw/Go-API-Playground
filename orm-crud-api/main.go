package main

import (
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
	"github.com/labstack/echo"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username	string
	DB_Password	string
	DB_Port		string
	DB_Host		string
	DB_Name		string
}

func InitDB() {

	config := Config{
		DB_Username	: "root",
		DB_Password	: "welcome12345",
		DB_Port		: "3306",
		DB_Host		: "localhost",
		DB_Name		: "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	ID 			int		`json:"id" form:"id"`
	Name		string	`json:"name" form:"name"`
	Email		string	`json:"email" form:"email"`
	Password	string	`json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

//get all users
func GetAllUsers(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{} {
		"message" : "success get all users",
		"users" : users,
	})
}

//get user by id
func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	var count int64
	DB.Model(&User{}).Where("id = ?", id).Count(&count)

	if count == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id not found",
		})
	} else {
		var user User
		if tx := DB.Find(&user, "id=?", id); tx.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "cannot fetch data",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    &user,
		})
	}
}

//create new user
func CreateUser(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{} {
		"message" : "success create new user",
		"user" : user,
	})
}

//delete user by id
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var user User
	if tx := DB.Find(&user, "id=?", id); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	if tx := DB.Delete(&user); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

//update user by id
func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var user User
	if tx := DB.Find(&user, "id=?", id); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	c.Bind(&user)
	if tx := DB.Save(&user); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

func main() {
	e := echo.New()

	//routing
	e.GET("/users", GetAllUsers)
	e.GET("/users/:id", GetUser)
	e.POST("/users", CreateUser)
	e.DELETE("/users/:id", DeleteUser)
	e.PUT("/users/:id", UpdateUser)

	//start the server
	e.Logger.Fatal(e.Start(":8080"))
}