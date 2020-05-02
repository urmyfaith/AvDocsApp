package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AddAdminEmail struct {
	gorm.Model
	Uniqueid	string 		`json: "uniqueid" `
	Flag		string 		`json: "flag"`
	Expirydate	time.Time	`json: "expiryDate"`
	ClinicmasterID	uint	`json: "clinicmasterID" sql:"index"`
}
