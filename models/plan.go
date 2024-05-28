package models

import "gorm.io/gorm"

type Plan struct {
	gorm.Model
	LectureName string `gorm:"unique;not null;foreigner:LectureName"`
	UserID      uint   `gorm:"unique;not null;foreigner:UserID"`
	StartTime   string
	EndTime     string
	Status      string
}
