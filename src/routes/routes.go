package routes

import (
	"net/http"

	"github.com/charoleizer/my-collection/src/domain"
	"github.com/labstack/echo"
)

func MapRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, domain.HelloSekai())
	})
}
