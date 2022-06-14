package routes

import (
	"net/http"

	"github.com/charoleizer/my-collection/src/domain"
	"github.com/labstack/echo"
)

func RunServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, domain.HelloSekai())
	})
	e.Logger.Fatal(e.Start(":1323"))
}
