package main

import (
	"projects/config"
	"projects/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDb()
	// config.InitPort()
	routes.New(e)
	e.Logger.Fatal(e.Start(":8000"))
	// e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
