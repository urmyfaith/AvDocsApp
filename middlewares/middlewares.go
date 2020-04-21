package middlewares

import (
	"AvDocsApp/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

//middleware customcontext
func MidCustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &model.CustomContext{c}
		return next(cc)
	}
}


func CorsPolicy()  middleware.CORSConfig {
	mid := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}
	return mid
}


