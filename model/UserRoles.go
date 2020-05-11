package model

import (
	"AvDocsApp/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Servicemaster struct {
	gorm.Model
	Name 		string		`json:"name"`
	Htmlname 	string		`json:"htmlname"`
	Paths 		string		`json:"paths"`
}


type Rightsservicemapper struct {
	gorm.Model
	Servicename				string			`json:"servicename"`
	Add 					bool			`json:"add"`
	Edit 					bool 			`json:"edit"`
	View					bool			`json:"view"`
	Delete 					bool			`json:"delete"`
	RightsmasterID			uint 			`sql:"index"`
}

type Right struct {
	Servicename				string			`json:"servicename"`
	Add 					bool			`json:"add"`
	Edit 					bool 			`json:"edit"`
	View					bool			`json:"view"`
	Delete 					bool			`json:"delete"`
}

type Rightsmaster struct {
	gorm.Model
	ClinicmasterID 			uint		 				`json:"clinicmaster_id"`
	Rightsname				string						`json:"rightsname"`
	Rightsservicemapper		[]Rightsservicemapper		`gorm :"FOREIGNKEY:RightsmasterID; json:"rightsservicemapper"`
}


func Allservicemasters() (servicemaster []Servicemaster) {
	dbs  := db.DbConn()
	dbs.Find(&servicemaster)
	fmt.Println(servicemaster)
	return
}

func Adminrightsonly() (servicemaster []Servicemaster) {
	dbs := db.DbConn()
	dbs.Not("name", []string{"addclinic", "addHot"}).Find(&servicemaster)
	return
}


func Getrights(id uint, clinicmasterid uint) (rightsmaster Rightsmaster, rightsservicemapper []Rightsservicemapper) {
	dbs := db.DbConn()
	defer dbs.Close()
	dbs.Where("id = ? and clinicmaster_id = ?", id, clinicmasterid).Find(&rightsmaster)
	dbs.Where("rightsmaster_id = ?", rightsmaster.ID).Find(&rightsservicemapper)
	return
}



