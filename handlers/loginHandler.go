package handlers

import (
	"AvDocsApp/model"
	//"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		u := new(model.Client)
		if err := c.Bind(u); err != nil {
			return err
		}
		counts := model.CheckLogin(u, db)
		fmt.Println("Username : ", u.Username)

		if counts == 0 {
			fmt.Println("unauthorised")
			return echo.ErrUnauthorized
		}
		// Throws unauthorized error
		//if u.Username != "sakibmulla@gmail.com" || u.Password != "12345678" {
		//	fmt.Println("unauthorised")
		//	return echo.ErrUnauthorized
		//}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = u.Username
		claims["role"] = "admin"
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
