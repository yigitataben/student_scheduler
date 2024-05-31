package models

import (
	"gorm.io/gorm"
	"time"
)

type Plan struct {
	gorm.Model
	LectureName string    `gorm:"unique;not null;foreigner:LectureName"`
	UserID      uint      `gorm:"unique;not null;foreigner:UserID"`
	StartTime   time.Time `gorm:"not null"`
	EndTime     time.Time `gorm:"not null"`
	Status      string    `gorm:"not null"`
}
