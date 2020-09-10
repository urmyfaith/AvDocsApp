package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Htmlhandlers() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var id string
		id = "<table border=" + "1" + ">"
		for i := 0; i < 10; i++ {
			id = id + "<tr><td>" + strconv.Itoa(i) + "</td></tr>"
		}
		id = id + "</table>"
		fmt.Println(id)
		return c.String(http.StatusOK, id)
	}
}
