package server

import (
	"github.com/charoleizer/my-collection/src/routes"
	"github.com/labstack/echo"
)

func RunServer() {
	e := echo.New()
	e.Static("/", "src/static/")
	e.File("/favicon.ico", "src/static/favicon.ico")
	e.File("/", "src/static/index.html")

	// e.Use(middleware.Logger())

	routes.MapRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
