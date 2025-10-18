package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name            string  `gorm:"size:255;not null"`
	Slug            string  `gorm:"size:255;not null"`
	Thumbnail       string  `gorm:"size:255;not null"`
	AboutBrand      string  `gorm:"type:text;not null"`
	DesignExecution string  `gorm:"type:text;not null"`
	OurWorkID       uint    `gorm:"not null"`
	Images          []Image `gorm:"foreignKey:ProjectID"`
}
