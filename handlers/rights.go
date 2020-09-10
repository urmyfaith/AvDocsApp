package handlers

import (
	"AvDocsApp/model"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
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


func GetRightsByClinicId() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.Param("id")
		fmt.Println("Clinic id is ======>> ", id)
		u, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, model.GetClinetwiserights(u))
	}
}

func Deleterights() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.Param("id")
		u, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return err
		}
		model.DeleterightByid(u)
		return
	}
}