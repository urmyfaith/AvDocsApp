package model

import (
	"AvDocsApp/db"
	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	Username string `json:"username" xml:"username" validate:"required,email"`
	Password string `json:"password" xml:"password" validate:"required"`
}

type Data struct {
	Id uint `json:"id"`
	Clinicid uint `json:"clinicid"`
}


func CheckLogin(u *Client) (loginmaster Loginmaster) {
	dbs := db.DbConn()
	defer dbs.Close()
	dbs.Where("emailid = ? AND password = ?", u.Username, u.Password).Find(&loginmaster)
	return
}

//func CheckLogin(u *Client) (loginmaster int) {
//	dbs := db.DbConn()
//	defer dbs.Close()
//	var c Client
//	dbs.Where("username = ? AND password = ?", u.Username, u.Password).Find(&c).Count(&loginmaster)
//	return
//}

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
