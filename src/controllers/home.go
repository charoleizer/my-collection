package controllers

import (
	"net/http"

	"github.com/charoleizer/my-collection/src/models"
	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	home := models.Home{}
	home.Reset()

	return c.Render(http.StatusOK, "home.html", home)
}
