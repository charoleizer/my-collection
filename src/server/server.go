package server

import (
	"github.com/charoleizer/my-collection/src/routes"
	"github.com/labstack/echo"
)

func RunServer() {
	e := echo.New()
	routes.MapRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
