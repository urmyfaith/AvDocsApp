package handlers

import (
	"AvDocsApp/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func Login() echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		u := new(model.Client)
		if err := c.Bind(u); err != nil {
			return err
		}

		fmt.Println("Username : ",u.Username)
		// Throws unauthorized error
		if u.Username != "sakibmulla@gmail.com" || u.Password != "12345678" {
			fmt.Println("unauthorised")
			return echo.ErrUnauthorized
		}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = u.Username
		claims["admin"] = true
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
