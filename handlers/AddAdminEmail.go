package handlers

import (
	"AvDocsApp/model"
	"github.com/labstack/echo"
	"net/http"
)

func AddAdminEmail() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.Param("id")
		return c.JSON(http.StatusOK, model.VerifyEmail(id))
	}
}
