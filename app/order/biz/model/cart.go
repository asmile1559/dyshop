package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint64 `gorm:"primaryKey"`
	UserId    uint64 `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CartItems []CartItem     `gorm:"foreignKey:CartId"`
}

type CartItem struct {
	ID        uint64 `gorm:"primaryKey"`
	CartId    uint64 `gorm:"index"`
	ProductId uint64
	Quantity  int
}
