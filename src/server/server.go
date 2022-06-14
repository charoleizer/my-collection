package server

import (
	"github.com/charoleizer/my-collection/src/routes"
	"github.com/charoleizer/my-collection/src/views"
	"github.com/labstack/echo"
)

func RunServer() {
	e := echo.New()
	e.Static("/", "src/views/public/")
	e.File("/favicon.ico", "src/views/public/favicon.ico")
	e.File("/", "src/views/public/index.html")

	// e.Use(middleware.Logger())

	views.RenderTemplate(e)
	routes.MapRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
