package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255;unique;not null"`
	Password string `gorm:"size:255;not null"`
}
