package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	Icon        string `json:"icon" gorm:"type:varchar(255);not null"`
}
