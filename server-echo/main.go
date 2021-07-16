package main

// go mod init github.com/<user>/<repo>
//go mod init <url>/<package>

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
)

func postHandler(c echo.Context) error {
	c.String(200, "hello world from echo: POST")
	return nil
}

func getHandler(c echo.Context) error {
	// c.String(200, "hello world from echo: GET")
	c.JSON(http.StatusOK, map[string]string{
		"name" : "john snow",
		"house" : "stark",
	})
	return nil
}

func triggerError(c echo.Context) error {
	return fmt.Errorf("ceritanya error")
}

func main() {
	e := echo.New()
	e.POST("/", postHandler)
	e.GET("/", getHandler)
	e.GET("/error", triggerError)
	e.GET("/users/:userId/order/:orderId", func (c echo.Context) error {
		userId := c.Param("userId")
		orderId := c.Param("orderId")
		return c.JSON(http.StatusOK, map[string]string {
			"user" : userId,
			"order" : orderId,
		})
	})
	e.GET("/users", func (c echo.Context) error {
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		return c.JSON(http.StatusOK, map[string]string {
			"limit" : limit,
			"offset" : offset,
		})
	})
	e.POST("/users", func (c echo.Context) error {
		name := c.FormValue("name")
		address := c.FormValue("address")
		return c.JSON(http.StatusOK, map[string]string {
			"name" : name,
			"address" : address,
		})
	})
	e.POST("/orders", func (c echo.Context) error {
		order := map[string]string{}
		c.Bind(&order)
		return c.JSON(http.StatusOK, order)
	})
	e.Start(":8080")
}