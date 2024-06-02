package models

import (
	"time"
)

type Lecture struct {
	ID          int `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt   time.Time
	LectureName string `gorm:"unique;not null"`
}
