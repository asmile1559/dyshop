package model

import "time"

// PaymentRecord 代表支付流水记录表
type PaymentRecord struct {
	ID            uint      `gorm:"primaryKey"`
	TransactionID string    `gorm:"transaction_id;type:varchar(64);not null"`    // 交易 ID
	Amount        string    `gorm:"not null"`               // 支付金额（字符串）
	Status        string    `gorm:"not null"`               // 交易状态
	CreatedAt     time.Time `gorm:"autoCreateTime"`         // 记录创建时间
}