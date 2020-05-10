package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Loginmaster struct {
	gorm.Model
	Firstname  string 		`json: "firstname" validate:"required"`
	Lastname string 		`json: "lastname" validate:"required"`
	Dob time.Time			`json: "dob" validate:"required"`
	RightsNo uint			`json: "rightsNo" `
	Homeaddress1 string 	`json: "homeaddress1" validate:"required"`
	Homeaddress2 string 	`json: "homeaddress2" validate:"required"`
	Homeaddress3 string 	`json: "homeaddress3" validate:"required"`
	Pincode  string			`json: "pincode" validate:"required"`
	Cityname string			`json: "cityname" validate: "required"`
	University string		`json: "university" validate: "required"`
	Degreename string		`json: "degreename" validate:"required"`
	Passoutyear string		`json: "passoutyear" validate:"required"`
	Policyaccept bool		`json: "policyaccept" validate:"required"`
	Password string			`json: "password" validate:"required"`
	Repassword string		`json: "repassword" validate:"required"`
	Oldpassword string		`json: "oldpassword" validate:"required"`
	Resetflag string		`json: "resetflag" `
	Emailid string			`gorm:"PRIMARY_KEY; NOT NULL, UNIQUE_INDEX" json: "emailid" validate:"required"`
	Contactno string		`gorm:"unique; NOT NULL, UNIQUE_INDEX" json: "contactno" validate:"required"`
	Reffcontno string		`json: "reffcontno" validate:"required"`
	Uniqueid string			`json: "uniqueid" validate:"required"`
	Statusflag string		`json: "statusflag"`
	ClinicmasterID	uint	`json: "clinicmasterID" sql:"index"`
}





