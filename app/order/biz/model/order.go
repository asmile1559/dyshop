package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           uint64 `gorm:"primaryKey"`
	UserId       uint64 `gorm:"index"`
	UserCurrency string
	Address      Address
	Email        string
	OrderItems   []OrderItem `gorm:"foreignKey:OrderId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Address struct {
	ID            uint64 `gorm:"primaryKey"`
	OrderId       uint64 `gorm:"uniqueIndex"`
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       string
}

type OrderItem struct {
	ID        uint64 `gorm:"primaryKey"`
	OrderId   uint64 `gorm:"Index"`
	ProductId uint64
	Quantity  int
	Cost      float64
}
