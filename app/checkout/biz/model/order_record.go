package model

import "time"

// 订单记录表
type OrderRecord struct {
	ID              uint      `gorm:"primaryKey"`
	OrderID         string    `gorm:"type:varchar(255);uniqueIndex"` // 订单号
	UserID          uint32    `gorm:"index"`       // 用户ID
	TransactionID   string    `gorm:"index"`       // 交易ID（用于查询）
	Recipient       string    // 收件人
	Phone           string    // 手机号
	Province        string    // 省
	City            string    // 市
	District        string    // 区
	Street          string    // 街道
	FullAddress     string    // 详细地址
	TotalQuantity   int       // 商品总数
	TotalPrice      float64   // 订单总价
	Postage         float64   // 邮费
	FinalPrice      float64   // 最终价格
	CreatedAt       time.Time // 创建时间
}

// 商品记录表
type OrderItem struct {
	ID            uint    `gorm:"primaryKey"`
	OrderID       string  `gorm:"index"`  // 关联订单号
	TransactionID string  `gorm:"index"`  // 交易ID（可用于查询）
	ProductID     string  // 商品ID
	ProductImg    string  // 商品图片
	ProductName   string  // 商品名称
	SpecName      string  // 规格名称
	SpecPrice     float64 // 规格价格
	Quantity      int     // 数量
	Postage       float64 // 邮费
	Currency      string  // 货币
}
