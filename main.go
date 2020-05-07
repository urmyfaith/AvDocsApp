package main

import (
	"AvDocsApp/db"
	"AvDocsApp/handlers"
	"AvDocsApp/middlewares"
	//"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	fmt.Println("Init")
	//creating tables
	//DbMonitor()
}

func DbMonitor() {
	dbs := db.DbConn()
	defer dbs.Close()
	//dbs.DropTableIfExists(&model.Loginmaster{})
	//dbs.DropTableIfExists(&model.Clinicmaster{}, &model.Telnumber{}, &model.AddAdminEmail{})
	//dbs.CreateTable(&model.Clinicmaster{}, &model.Telnumber{}, &model.AddAdminEmail{})
	//dbs.CreateTable(&model.Loginmaster{})
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	fmt.Println(i)
	return cv.validator.Struct(i)
}

func main() {
	// features.MailCompose()
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
	e.POST("/login", handlers.Login())

	//addAdminEmail
	e.GET("/addadmin/:id", handlers.AddAdminEmail())

	//addAdmin
	e.POST("/saveadmin", handlers.SaveAdmin())

	// Restricted group
	r := e.Group("")

	r.Use(middleware.JWT([]byte("secret")))

	//login
	r.GET("/", handlers.Dashboard())

	//addclinic
	r.POST("/addclinic", handlers.AddClinic())

	//start server
	e.Logger.Fatal(e.Start(":8000"))
}

