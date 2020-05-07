package handlers

import (
	"AvDocsApp/db"
	"AvDocsApp/model"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func SaveAdmin() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		loginmaster := new(model.Loginmaster)
		if err := c.Bind(loginmaster); err != nil {
			fmt.Print("badsssssssssss")
			return err
		}
		if err = c.Validate(loginmaster); err != nil {
			return err
		}
		s:= model.UpdateFlag(loginmaster.Uniqueid)
		fmt.Println("s =>>>>>>>>>> ", s)
		if ( s.Flag == "B") {
			loginmaster.RightsNo = 100
			loginmaster.Resetflag = "S"  // set s and for reset R
			loginmaster.Statusflag = "A" // A = Active and L = Lock and  B = Block
		}
			fmt.Println("===============>>>>>>>>>>", loginmaster)
			dbs := db.DbConn()
			defer dbs.Close()
			dbs.Create(loginmaster)
		return c.JSON(http.StatusOK, loginmaster)
	}
}
