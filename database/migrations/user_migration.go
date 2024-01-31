package migrations

import (
	"crud-example-go/database/models"

	"github.com/jinzhu/gorm"
)

func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
