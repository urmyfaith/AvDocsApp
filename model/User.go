package model

type Client struct {
	Username string `json:"username" xml:"username" validate:"required,email"`
	Password string `json:"password" xml:"password" validate:"required"`
}


