package model

import (

)

type User struct {	
	UserID    int64    `gorm:"primary_key"`
	Email     string    `gorm:"unique;not null" `
	Password  string    `gorm:"not null" binding:"required"`
}
