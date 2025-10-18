package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	URL       string `gorm:"size:255;not null"`
	ProjectID uint   `gorm:"not null"`
}
