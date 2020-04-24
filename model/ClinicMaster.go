package model

import "time"

type ClinicMaster struct {
	ClinicId    int64     `json:"clinicId" validate:"required"`
	ClinicName  string    `json:"clinicName" validate:"required"`
	ContactNo   []ClinicContact     `json:"contactNo" `
	Cityname   	string	  `json:"cityname" validate:"required"`
	Address     string    `json:"address"`
	Pincode     int       `json:"pincode" validate:"required"`
	District    string    `json:"district" validate:"required"`
	Country     string    `json:"country"`
	CreatedDate time.Time `json:"createdDate"`
	CreatedBy   int64     `json:"createdBy" validate:"required"`
}

type ClinicContact struct {
	ClinicId    int64     `json:"clinicId" validate:"required"`
	ClinicName  string    `json:"clinicName" validate:"required"`
	ContactNo   int64     `json:"contactNo" validate:"required"`
}
