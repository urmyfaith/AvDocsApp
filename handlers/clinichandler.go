package handlers

import (
	"AvDocsApp/db"
	"AvDocsApp/features"
	"AvDocsApp/model"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func AddClinic() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		fmt.Println("reqest come inside")
		clinic := new(model.Clinicmaster)
		if err := c.Bind(clinic); err != nil {
			fmt.Print("badsssssssssss")
			return err
		}
		if err = c.Validate(clinic); err != nil {
			return err
		}
		dbs := db.DbConn()
		defer dbs.Close()
		dbs.Create(clinic)
		fmt.Println("id is ",clinic.ID)
		addAdmin := &model.AddAdminEmail{Uniqueid: features.Encryption(clinic.ClinicName+""+clinic.Emailid+""+clinic.Address), Flag: "A", Expirydate: time.Now(), ClinicmasterID: clinic.ID}
		features.MailCompose(addAdmin.Uniqueid, clinic.Emailid)
		dbs.Create(addAdmin)
		return c.JSON(http.StatusOK, clinic)
	}
}