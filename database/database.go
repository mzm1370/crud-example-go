package database

import (
	"crud-example-go/database/migrayions"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/splite"
)

var DB *gorm.db

func init() {
	var err error
	DB, err = gorm.Open("splite3", "store")

	if err != nil {
		panic("Failed to connect to data")
	}

	migrayions.MigrateUser(DB)

}
