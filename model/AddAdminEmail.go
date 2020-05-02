package model

import (
	"AvDocsApp/db"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type AddAdminEmail struct {
	gorm.Model
	Uniqueid	string 		`json: "uniqueid" `
	Flag		string 		`json: "flag"`
	Expirydate	time.Time	`json: "expiryDate"`
	Role 		string		`json: "role"`
	ClinicmasterID	uint	`json: "clinicmasterID" sql:"index"`
}


func VerifyEmail(id string) (addAdminEmail AddAdminEmail) {
	dbs := db.DbConn()
	defer dbs.Close()
	fmt.Println(id)
	// Get first matched record
	err := dbs.Where("uniqueid = ?", id).Find(&addAdminEmail)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(&addAdminEmail)
	return
}