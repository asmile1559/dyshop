package model

import (
	"time"
)

// Address 结构体，用于描述地址信息
type Address struct {
	AddressId   string `gorm:"primaryKey,column:address_id"` // 对应 proto 中的 AddressId 字段
	Recipient   string `gorm:"column:recipient"`             // 新增字段
	Phone       string `gorm:"column:phone"`                 // 新增字段
	Province    string `gorm:"column:province"`              // 新增字段
	City        string `gorm:"column:city"`
	District    string `gorm:"column:district"`     // 新增字段
	Street      string `gorm:"column:street"`       // 新增字段
	FullAddress string `gorm:"column:full_address"` // 新增字段
}

// ProductSpec 结构体，用于描述商品规格
type ProductSpec struct {
	Name  string  `gorm:"primaryKey,column:name"`
	Price float64 `gorm:"column:price"`
}

// Product 结构体，用于描述商品信息
type Product struct {
	ProductID   uint64      `gorm:"primaryKey,column:product_id"`
	ProductImg  string      `gorm:"column:product_img"`
	ProductName string      `gorm:"column:product_name"`
	ProductSpec ProductSpec `gorm:"embedded"`
	Quantity    int         `gorm:"column:quantity"`
	Currency    string      `gorm:"column:currency"`
	Postage     float64     `gorm:"column:postage"`
}

//// UserInfo 结构体，用于描述用户信息
//type UserInfo struct {
//	Name string `gorm:"column:name"`
//}

// Order 结构体，用于描述订单信息
type Order struct {
	OrderID         uint64    `gorm:"primaryKey"`
	UserId          uint32    `gorm:"index"` // 修改为 uint32 以匹配 proto 中的 user_id 类型
	UserCurrency    string    `gorm:"column:user_currency"`
	Address         Address   `gorm:"embedded"` // 嵌入式 Address
	Email           string    `gorm:"column:email"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	OrderPrice      float64   `gorm:"column:order_price"`
	OrderPostage    float64   `gorm:"column:order_postage"`
	OrderDiscount   float64   `gorm:"column:order_discount"`
	OrderFinalPrice float64   `gorm:"column:order_final_price"`
	Products        []Product `gorm:"foreignKey:ProductID"`
}
