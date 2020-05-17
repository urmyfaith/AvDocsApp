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
	gorm.Model
	Servicename				string			`json:"servicename"`
	Add 					bool			`json:"add"`
	Edit 					bool 			`json:"edit"`
	View					bool			`json:"view"`
	Delete 					bool			`json:"delete"`
	Comments				string 			`json:"comments"`
	Htmltag					string			`json:"htmltag"`
}

type Rightdata struct {
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

type RightsApi struct {
	Id  			uint		`json:"id" validate:"required"`
	Clinicid		uint		`json:"clinicid" validate:"required"`
	Rightsname 		string		`json:"rightsname" validate:"required"`
	Rightdata		[]Right		`json:"rightdata" validate:"required"`
}


func SaveRights(rr *RightsApi) (right []Right) {
	var rightsmaster Rightsmaster
	rightsmaster.ClinicmasterID = rr.Clinicid
	rightsmaster.Rightsname = rr.Rightsname
	for _, v := range rr.Rightdata {
		rightsmaster.Rightsservicemapper = append(rightsmaster.Rightsservicemapper, Rightsservicemapper{Servicename: v.Servicename, Add: v.Add, Edit: v.Edit, View: v.View, Delete: v.Delete})
	}
	dbs  := db.DbConn()
	defer dbs.Close()
	dbs.Create(&rightsmaster)
	fmt.Println("rights saved ",rightsmaster)
	var rightsservicemapper []Rightsservicemapper
	if rightsmaster.ID != 0 {
		dbs.Where("rightsmaster_id = ?", rightsmaster.ID).Find(&rightsservicemapper)
		for _, v := range rightsservicemapper {
			right = append(right, Right{Servicename: v.Servicename, Htmltag: "ok", Comments: "ok", Delete: v.Delete, Add: v.Add, Edit: v.Edit, View: v.View})
		}
	}
	return
}



func Allservicemasters() (servicemaster []Servicemaster) {
	dbs  := db.DbConn()
	defer dbs.Close()
	dbs.Find(&servicemaster)
	fmt.Println(servicemaster)
	return
}

func Adminrightsonly() (servicemaster []Servicemaster) {
	dbs := db.DbConn()
	dbs.Not("name", []string{"clinic_managements", "addHot"}).Find(&servicemaster)
	return
}

func Getrights(id uint, clinicmasterid uint) (rightsmaster Rightsmaster, rightsservicemapper []Rightsservicemapper) {
	dbs := db.DbConn()
	defer dbs.Close()
	dbs.Where("id = ? and clinicmaster_id = ?", id, clinicmasterid).Find(&rightsmaster)
	dbs.Where("rightsmaster_id = ?", rightsmaster.ID).Find(&rightsservicemapper)
	return
}

func GetAllRights() (right []Right) {
	dbs := db.DbConn()
	defer dbs.Close()
	dbs.Not("servicename", []string{"clinic_managements", "addHot"}).Find(&right)
	return
}

func GetClinetwiserights(id uint64) (rightsmaster []Rightsmaster) {
	dbs := db.DbConn()
	defer dbs.Close()
	dbs.Where("clinicmaster_id = ? ", id).Find(&rightsmaster)
	return
}



