package models

import "time"

type User struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
}
