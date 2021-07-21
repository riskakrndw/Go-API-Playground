package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        int    `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}

var db *gorm.DB

func main() {
	connectionString := "root:welcome12345@tcp(localhost:3306)/orm?parseTime=true"
	var err error
	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Book{})
	fmt.Println(db)

	e := echo.New()
	e.GET("/books", GetBooks)
	e.GET("/books/:id", GetOneBook)
	e.POST("/books", CreateBook)
	e.PUT("/books/:id", UpdateBook)
	e.DELETE("/books/:id", DeleteBook)
	e.Start(":8080")
}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var book Book
	if tx := db.Find(&book, "id=?", id); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	if tx := db.Delete(&book); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    book,
	})
}

func UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var book Book
	if tx := db.Find(&book, "id=?", id); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	c.Bind(&book)
	if tx := db.Save(&book); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    book,
	})
}

func GetOneBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var book Book
	if tx := db.Find(&book, "id=?", id); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    book,
	})
}

func CreateBook(c echo.Context) error {
	book := Book{}
	c.Bind(&book)
	if tx := db.Save(&book); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    book,
	})
}

func GetBooks(c echo.Context) error {
	books := []Book{}
	if tx := db.Find(&books); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    books,
	})
}
