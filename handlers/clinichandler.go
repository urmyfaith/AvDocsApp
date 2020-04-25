package handlers

import (
	"AvDocsApp/model"
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func AddClinic(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		fmt.Println("reqest come inside")
		clinic := new(model.ClinicMaster)
		if err := c.Bind(clinic); err != nil {
			fmt.Print("badsssssssssss")
			return err
		}
		return c.JSON(http.StatusOK, clinic)
	}
}