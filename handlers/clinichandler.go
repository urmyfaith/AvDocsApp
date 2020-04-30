package handlers

import (
	"AvDocsApp/model"
	//"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"github.com/jinzhu/gorm"
)

func AddClinic(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		fmt.Println("reqest come inside")
		clinic := new(model.Clinicmaster)
		if err := c.Bind(clinic); err != nil {
			fmt.Print("badsssssssssss")
			return err
		}
		defer db.Close()
		db.Create(clinic)
		return c.JSON(http.StatusOK, clinic)
	}
}