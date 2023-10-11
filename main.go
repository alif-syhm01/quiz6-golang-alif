package main

import (
	"echo-golang-quiz6/configs"
	"echo-golang-quiz6/routes"

	"github.com/labstack/echo/v4"
)

func init() {
	configs.InitDB()
	configs.InitialMigration()
}

func main() {
	// create a new echo instance
	e := echo.New()

	// Routing
	routes.InitRoute(e)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
