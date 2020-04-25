package main

import (
	"AvDocsApp/db"
	"AvDocsApp/handlers"
	"AvDocsApp/middlewares"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var MysqlPool *sql.DB

func init() {
	fmt.Println("Init")
	MysqlPool = db.DbConn()
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// Echo instance
	e := echo.New()

	// middlewares
	e.Use(middlewares.MidCustomContext)

	//middlewares custom cors
	e.Use(middleware.CORSWithConfig(middlewares.CorsPolicy()))

	//request model validator.
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", handlers.Login(MysqlPool))

	// Restricted group
	r := e.Group("")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/", handlers.Dashboard())
	r.POST("/addclinic", handlers.AddClinic(MysqlPool))

	//start server
	e.Logger.Fatal(e.Start(":8000"))
}
