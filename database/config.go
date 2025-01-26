package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Ini configurationnya buat connect ke database.
func ConnectDatabase() {
	var err error

	//Ini dari docs gorm.io buat connect ke database.
	const MySQL = "root:@tcp(127.0.0.1:3306)/crudvalorant?charset=utf8mb4&parseTime=True&loc=Local"
	DSN := MySQL

	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Can't connect to DB")
	} else {
		fmt.Println("Connected to database")
	}
}
