package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"unique;not null" json:"email"`
	Address  string `json:"address"`
	Tel      string `json:"tel"`
}
