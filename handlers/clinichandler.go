package handlers

import (
	"AvDocsApp/db"
	"AvDocsApp/model"
	//"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func AddClinic() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		fmt.Println("reqest come inside")
		clinic := new(model.Clinicmaster)
		if err := c.Bind(clinic); err != nil {
			fmt.Print("badsssssssssss")
			return err
		}
		dbs := db.DbConn()
		defer dbs.Close()
		dbs.Create(clinic)
		return c.JSON(http.StatusOK, clinic)
	}
}