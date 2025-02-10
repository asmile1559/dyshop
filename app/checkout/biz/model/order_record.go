package model

import "time"

// Order 订单数据模型
type OrderRecord struct {
	ID            uint      `gorm:"primaryKey"`
	OrderID       string    `gorm:"type:varchar(255);uniqueIndex"`
	UserID        uint32    `gorm:"not null"`
	Amount        float64   `gorm:"not null"`
	TransactionID string    `gorm:"not null"`
	Status        string    `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}
