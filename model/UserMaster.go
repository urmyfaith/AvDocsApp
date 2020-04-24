package model

import (
	"time"
)

type UserMasters struct {
	ClinicId    int64     `json:"clinicId"`
	Id          int64     `json:"id"`
	FirstName   string    `json:"firstName" validate:"required"`
	LastName    string    `json:"lastName" validate:"required"`
	Dob         time.Time `json:"dob"`
	Email       string    `json:"email" validate:"required,email"`
	ContactNo   int32     `json:"contactNo" validate:"required"`
	StatusFlag  string    `json:"statusFlag" validate:"required"`
	CreatedDate time.Time `json:"createdDate" `
	ExpiryDate  time.Time `json:"expiryDate" `
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	SessionFlag string    `json:"sessionFlag"`
	Role        string    `json:"role" validate:"required"`
	CreatedBy   int64     `json:"createdBy" validate:"required"`
}
