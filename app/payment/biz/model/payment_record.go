package model

import "time"

// PaymentRecord 代表支付流水记录表
type PaymentRecord struct {
	ID            uint      `gorm:"primaryKey"`
	OrderID       string    `gorm:"column:order_id;type:varchar(64);not null"`
	UserID        uint32    `gorm:"column:user_id;not null"`
	Amount        float32   `gorm:"column:amount;not null"`
	TransactionID string    `gorm:"column:transaction_id;type:varchar(64);not null"`
	Status        string    `gorm:"column:status;type:varchar(32);not null"` // SUCCESS, FAILED
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
}