package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"default:null"`
	Completed   bool   `gorm:"not null;default:false"`
}
