package models

import (
	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	LectureID uint    `gorm:"unique;not null"`
	Lecture   Lecture `gorm:"foreignKey:LectureID"`
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	StartTime string  `gorm:"not null"`
	EndTime   string  `gorm:"not null"`
	Status    string  `gorm:"not null"`
}
