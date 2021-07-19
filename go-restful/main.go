package main

import (
	"alta/training/config"
	"alta/training/middlewares"
	"alta/training/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDb()
	// config.InitPort()
	middlewares.LogMiddlewares(e)
	routes.New(e)
	// e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
	e.Logger.Fatal(e.Start(":8000"))
}
