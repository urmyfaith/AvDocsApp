package model

import "time"

type ClinicMaster struct {
	ClinicId    int64     `json:"clinicId" validate:"required"`
	ClinicName  string    `json:"clinicName" validate:"required"`
	ContactNo   int32     `json:"contactNo" validate:"required"`
	Address     string    `json:"address" validate:""`
	Pincode     int       `json:"pincode" validate:"required"`
	District    string    `json:"district" validate:"required"`
	Country     string    `json:"country" validate:""`
	CreatedDate time.Time `json:"country"`
	CreatedBy   int64     `json:"createdBy" validate:"required"`
}
