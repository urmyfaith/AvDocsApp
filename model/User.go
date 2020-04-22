package model

import (
	"database/sql"
	"fmt"
)

type Client struct {
	Username string `json:"username" xml:"username" validate:"required,email"`
	Password string `json:"password" xml:"password" validate:"required"`
}

func CheckLogin(u *Client, db *sql.DB) int {
	var counts int
	selDB, err := db.Query("select count(*) from Login_master where email= ? and password = ?", u.Username, u.Password)
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		fmt.Println("scan ")
		err = selDB.Scan(&counts)
		if err != nil {
			panic(err.Error())
		}
	}
	return counts
}
