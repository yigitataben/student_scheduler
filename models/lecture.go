package models

import "gorm.io/gorm"

type Lecture struct {
	gorm.Model
	LectureName string `gorm:"unique;not null"`
	LectureID   uint   `gorm:"unique;not null"`
}