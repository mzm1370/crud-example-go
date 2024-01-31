package database

import (
	"crud-example-go/database/migrations"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("splite3", "store")

	if err != nil {
		panic("Failed to connect to data")
	}

	migrations.MigrateUser(DB)

}
