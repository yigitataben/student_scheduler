package models

import "gorm.io/gorm"

type Plan struct {
	gorm.Model
	LectureName string `gorm:"unique;not null;foreigner:LectureName"`
	UserID      uint   `gorm:"unique;not null;foreigner:UserID"`
	StartTime   string `gorm:"not null"`
	EndTime     string `gorm:"not null"`
	Status      string `gorm:"not null"`
}
