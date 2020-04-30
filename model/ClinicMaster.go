package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Clinicmaster struct {
	gorm.Model
	ClinicName  string    `json:"clinicName" validate:"required"`
	ContactNo   []Telnumber `gorm:"foreignkey:ClinicmasterID" json:"contactNo" `
	Cityname   	string	  `json:"cityname" validate:"required"`
	Address     string    `json:"address"`
	Pincode     int       `json:"pincode" validate:"required"`
	District    string    `json:"district" validate:"required"`
	Country     string    `json:"country"`
	CreatedDate time.Time `json:"createdDate"`
	CreatedBy   int64     `json:"createdBy" validate:"required"`
}

type Telnumber struct {
	gorm.Model
	Telenumber    	int64		`json:"telenumber" validate:"required"`
	ClinicmasterID	uint		`sql:"index"`
}
