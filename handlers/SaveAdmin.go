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
		if (s.Role == "rootadmin") {
			rightmaster := &model.Rightsmaster{
				ClinicmasterID: s.ClinicmasterID,
				Rightsname: "rootadmin",
				Rightsservicemapper: RightsDecider(s.Role),
			}
			dbs := db.DbConn()
			fmt.Println("before Rights Mapper ======> ",rightmaster)
			dbs.Create(&rightmaster)
			dbs.Close()
			fmt.Println("after Rights Mapper ======> ",rightmaster)
			if ( s.Flag == "B") {
				loginmaster.RightsNo = rightmaster.ID
				loginmaster.Resetflag = "S"  // set s and for reset R
				loginmaster.Statusflag = "A" // A = Active and L = Lock and  B = Block
				loginmaster.ClinicmasterID = s.ClinicmasterID
			}
		} else if (s.Role == "admin") {
			rightmaster := &model.Rightsmaster{
				ClinicmasterID: s.ClinicmasterID,
				Rightsname: "admin",
				Rightsservicemapper: RightsDecider(s.Role),
			}
			dbs := db.DbConn()
			fmt.Println("before Rights Mapper ======> ",rightmaster)
			dbs.Create(&rightmaster)
			dbs.Close()
			fmt.Println("after Rights Mapper ======> ",rightmaster)
			if ( s.Flag == "B") {
				loginmaster.RightsNo = rightmaster.ID
				loginmaster.Resetflag = "S"  // set s and for reset R
				loginmaster.Statusflag = "A" // A = Active and L = Lock and  B = Block
				loginmaster.ClinicmasterID = s.ClinicmasterID
			}
		}
			fmt.Println("===============>>>>>>>>>>", loginmaster)
			dbs := db.DbConn()
			dbs.Create(loginmaster)
			dbs.Close()
		return c.JSON(http.StatusOK, loginmaster)
	}
}


func RightsDecider(role string) (Rightsservicemapper []model.Rightsservicemapper){
	if (role == "rootadmin") {
		allsermasters:=model.Allservicemasters()
		for _, f := range allsermasters {
			Rightsservicemapper = append(Rightsservicemapper, model.Rightsservicemapper{Servicename:f.Name, Add: true, Edit: true, View: true, Delete: true})
		}
	} else if (role == "admin") {
		allsermasters:=model.Adminrightsonly()
		for _, f := range allsermasters {
			Rightsservicemapper = append(Rightsservicemapper, model.Rightsservicemapper{Servicename:f.Name, Add: true, Edit: true, View: true, Delete: true})
		}
	}
	return
}


