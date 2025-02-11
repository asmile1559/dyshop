package model

import (
	"time"

)

type User struct {	
	UserID    uint32    `gorm:"primary_key"`
	Email     string    `gorm:"unique"`
	Password  string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
