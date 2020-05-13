package handlers

import (
	"AvDocsApp/model"
	"github.com/labstack/echo"
	"net/http"
)

func Sendrights() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, model.GetAllRights())
	}
}