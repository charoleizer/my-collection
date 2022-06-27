package controllers

import (
	"net/http"

	"github.com/charoleizer/my-collection/src/domain"
	"github.com/charoleizer/my-collection/src/models"
	"github.com/charoleizer/my-collection/src/repositories"
	"github.com/labstack/echo"
)

var f *repositories.Figures

func init() {
	f = &repositories.Figures{}
}

func Save(ctx echo.Context) error {

	// Loads the figure from the argument
	if ctx.Request().Method == "GET" {
		f.Figures = &models.Figures{}
		f.Figures.ID = "1"
		f.Figures.Name = "Luffy"
	}

	if ctx.Request().Method == "POST" {
		return ctx.String(http.StatusNotImplemented, "POST Method is comming soon")
	}

	// Checks if the figure is Luffy
	if !domain.IsLuffy(*f.Figures) {
		return ctx.String(http.StatusPreconditionFailed, "the figure is not Luffy")
	}

	// Saves the figure
	err := f.Save()
	if err != nil {
		return ctx.String(http.StatusExpectationFailed, err.Error())
	}

	return ctx.String(http.StatusOK, "Done")
}
