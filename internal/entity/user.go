package entity

import (
	"time"
)

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;size:36;not null"`
	Pass string `gorm:"size:36;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
