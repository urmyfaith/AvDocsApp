package handlers

import (
	"AvDocsApp/model"
	//"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		u := new(model.Client)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err = c.Validate(u); err != nil {
			return err
		}
		counts := model.CheckLogin(u)
		fmt.Println("Username : ", u.Username)

		if counts.Emailid != u.Username {
			fmt.Println("unauthorised")
			return echo.ErrUnauthorized
		}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		s, d := model.Getrights(counts.RightsNo, counts.ClinicmasterID)
		fmt.Println("rights name =======> ", s)
		fmt.Println("rights number =======> ", d)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = u.Username
		claims["role"] = "admin"
		claims["user"] = model.Data{Id: counts.ID, Clinicid: counts.ClinicmasterID}
		claims["rights"] = Rights(d)
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
}

func Rights(rights []model.Rightsservicemapper) (right []model.Rightdata) {
	for _, f := range rights {
		right = append( right, model.Rightdata{Servicename: f.Servicename, Add: f.Add, Edit: f.Edit, Delete: f.Delete, View: f.View})
	}
	return
}


//func Login() echo.HandlerFunc {
//	return func(c echo.Context) (err error) {
//		u := new(model.Client)
//		if err := c.Bind(u); err != nil {
//			return err
//		}
//		if err = c.Validate(u); err != nil {
//			return err
//		}
//		counts := model.CheckLogin(u)
//		fmt.Println("Username : ", u.Username)
//
//		if counts == 0 {
//			fmt.Println("unauthorised")
//			return echo.ErrUnauthorized
//		}
//
//		// Create token
//		token := jwt.New(jwt.SigningMethodHS256)
//
//		// Set claims
//		claims := token.Claims.(jwt.MapClaims)
//		claims["name"] = u.Username
//		claims["role"] = "admin"
//		claims["user"] = counts
//		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
//
//		// Generate encoded token and send it as response.
//		t, err := token.SignedString([]byte("secret"))
//		if err != nil {
//			return err
//		}
//
//		return c.JSON(http.StatusOK, map[string]string{
//			"token": t,
//		})
//	}
//}
