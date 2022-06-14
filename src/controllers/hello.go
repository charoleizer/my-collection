package controllers

import (
	"net/http"

	"github.com/charoleizer/my-collection/src/domain"
	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, domain.HelloSekai())
}
