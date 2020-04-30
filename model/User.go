package model

import (
	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	Username string `json:"username" xml:"username" validate:"required,email"`
	Password string `json:"password" xml:"password" validate:"required"`
}

func CheckLogin(u *Client, db *gorm.DB) int {
	defer db.Close()
	var counts int
	var clients Client
	db.Where("username = ? AND password = ?", u.Username, u.Password).Find(&clients).Count(&counts)
	return counts
}



//func CheckLogin(u *Client, db *sql.DB) int {
//	var counts int
//	selDB, err := db.Query("select count(*) from Login_master where email= ? and password = ?", u.Username, u.Password)
//	if err != nil {
//		panic(err.Error())
//	}
//	for selDB.Next() {
//		fmt.Println("scan ")
//		err = selDB.Scan(&counts)
//		if err != nil {
//			panic(err.Error())
//		}
//	}
//	return counts
//}
