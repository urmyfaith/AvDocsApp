package db

import (
	//"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/go-sql-driver/mysql"
)


func DbConn() (dbs *gorm.DB) {
		dbDriver := "mysql"
		dbUser := "skm"
		dbPass := "SakibMulla12!@"
		dbName := "gomysql"
		dbs, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp(166.62.29.42:3306)/"+dbName+"?parseTime=true")
		if err != nil {
			panic(err.Error())
		}
	return dbs
}

//func DbTables(dbs *gorm.DB) {
//	defer dbs.Close()
//	//dbs.DropTableIfExists(&model.Clinicmaster{}, &model.Telnumber{})
//	//dbs.CreateTable(&model.Client{})
//	//dbs.CreateTable(&model.Clinicmaster{}, &model.Telnumber{})
//	//dbs.AutoMigrate(&model.Client{})
//	//dbs.AutoMigrate(&model.Telnumber{})
//	//dbs.Create(&model.Client{Username: "sakib@gmail.com", Password: "12345678"})
//}


//func DbConn() (db *sql.DB) {
//	dbDriver := "mysql"
//	dbUser := "skm"
//	dbPass := "SakibMulla12!@"
//	dbName := "gomysql"
//	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(166.62.29.42:3306)/"+dbName)
//	if err != nil {
//		panic(err.Error())
//	}
//	db.SetConnMaxLifetime(500)
//	db.SetMaxIdleConns(50)
//	db.SetMaxOpenConns(10)
//	db.Stats()
//	return db
//}





