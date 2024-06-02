package models

import "time"

type Plan struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	LectureID int     `gorm:"unique;not null"`
	Lecture   Lecture `gorm:"foreignKey:LectureID"`
	UserID    int     `gorm:"unique;not null"`
	User      User    `gorm:"foreignKey:UserID"`
	StartTime string  `gorm:"not null"`
	EndTime   string  `gorm:"not null"`
	Status    string  `gorm:"not null"`
}
