package routes

import (
	"github.com/charoleizer/my-collection/src/controllers"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Echo) {
	e.GET("/hello", controllers.Hello)

	// TODO - Change to POST
	e.GET("/figures", controllers.Figures)

}
