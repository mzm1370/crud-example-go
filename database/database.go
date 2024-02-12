package database

import (
	"crud-example-go/database/migrations"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/moochin?charset=utf8&parseTime=True&loc=Local")
	println(err)
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to data")
	}

	migrations.MigrateUser(DB)

}
