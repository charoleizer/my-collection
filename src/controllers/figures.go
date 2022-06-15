package controllers

import (
	"net/http"

	"github.com/charoleizer/my-collection/src/domain"
	"github.com/charoleizer/my-collection/src/repositories"
	"github.com/labstack/echo"
)

func Save(f repositories.IFigures) error {
	err := f.Save()
	if err != nil {
		return err
	}
	return nil
}

func Figures(c echo.Context) error {
	figure := &repositories.Figures{}
	figure.Figures = domain.HelloFigures()

	err := Save(figure)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	return c.String(http.StatusOK, "Done")

}
