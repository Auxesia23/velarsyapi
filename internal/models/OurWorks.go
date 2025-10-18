package models

import "gorm.io/gorm"

type OurWork struct {
	gorm.Model
	Title    string    `gorm:"size:255"`
	Slug     string    `gorm:"size:255"`
	Image    string    `gorm:"size:255"`
	Projects []Project `gorm:"foreignKey:OurWorkID"`
}
