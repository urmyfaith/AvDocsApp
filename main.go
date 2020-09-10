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
	DbMonitor()
}

func DbMonitor() {
	dbs := db.DbConn()
	defer dbs.Close()
	//dbs.DropTableIfExists(model.Loginmaster{},model.Servicemaster{}, model.Clinicmaster{}, model.Telnumber{},model.AddAdminEmail{},model.Rightsservicemapper{}, model.Rightsmaster{})
	//dbs.CreateTable(model.Loginmaster{},model.Servicemaster{}, model.Clinicmaster{}, model.Telnumber{},model.AddAdminEmail{}, model.Rightsservicemapper{}, model.Rightsmaster{})
	//dbs.Create(model.Servicemaster{Name: "addclinic", Htmlname: "Add Clinic", Paths: "/ClinicManagement"})
	//dbs.Create(model.Servicemaster{Name: "adduser", Htmlname: "Add user", Paths: "/UserManagement"})
	//dbs.Create(model.Servicemaster{Name: "todayappointment", Htmlname: "Today Appointment", Paths: "/opdTodayappointment"})
	//dbs.Create(model.Servicemaster{Name: "newappointment", Htmlname: "New Opd Appointment", Paths: "/opdnewappointment"})
	//dbs.Create(model.Servicemaster{Name: "opdappointmenthistory", Htmlname: "Opd Appointment History", Paths: "/opdhisappointment"})
	//dbs.Create(model.Servicemaster{Name: "rights", Htmlname: "Rights", Paths: "/rightsmanagement"})
	//dbs.Create(model.Rightsservicemapper{Servicename: "rights", Add: true, Edit: true, View: true, Delete: true, RightsmasterID: 1})
	//dbs.Create(model.Rightsservicemapper{Servicename: "rights", View: true, Add: true, Edit: false, Delete: false, RightsmasterID: 2})
	//dbs.CreateTable(model.Right{})
	//dbs.Create(model.Right{Servicename: "clinic_managements", Add: false, Edit: false, View: false, Delete: false, Comments: "Clinic Management", Htmltag: "Clinic Management"})
	//dbs.Create(model.Right{Servicename: "user_managements", Add: false, Edit: false, View: false, Delete: false, Comments: "User Management", Htmltag: "User Management"})
	//dbs.Create(model.Right{Servicename: "rights", Add: false, Edit: false, View: false, Delete: false, Comments: "Rights Management", Htmltag: "Rights Management"})
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

	//send all rights
	r.GET("/sendrights", handlers.Sendrights())

	//get new rights from frontend
	r.POST("/saverights", handlers.Saverights())

	//get clinic all rights by clinic_id
	r.GET("/getAllRightsbyId/:id", handlers.GetRightsByClinicId())

	//delete rights by id
	r.DELETE("/deleteright/:id", handlers.Deleterights())

	//html test
	r.GET("/html", handlers.Htmlhandlers())

	//start server
	e.Logger.Fatal(e.Start(":8000"))
}
