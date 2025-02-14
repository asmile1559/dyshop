package model

import (

)

type User struct {	
	UserID    uint32    `gorm:"primary_key"`
	Email     string    `gorm:"unique"`
	Password  string    `gorm:"not null"`
}
