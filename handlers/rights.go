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


func Saverights() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		u := new(model.RightsApi)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err = c.Validate(u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, model.SaveRights(u))
	}
}