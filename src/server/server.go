package server

import (
	"github.com/charoleizer/my-collection/src/routes"
	"github.com/charoleizer/my-collection/src/views"
	"github.com/labstack/echo"
)

func RunServer() {
	e := echo.New()

	views.RenderTemplate(e)
	routes.MapRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
