package model

import "gorm.io/gorm"

// Address 结构体，用于描述地址信息
type Address struct {
	gorm.Model
	UserId      int64  `gorm:"column:uid"`
	Phone       string `gorm:"column:phone"`     // 新增字段
	Recipient   string `gorm:"column:recipient"` // 新增字段
	Province    string `gorm:"column:province"`  // 新增字段
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

// Order 结构体，用于描述订单信息
type Order struct {
	gorm.Model
	UserId     int64   `gorm:"column:user_id"`
	AddressId  uint32  `gorm:"column:address_id"`
	Price      float64 `gorm:"column:price"`
	ProductIDs string  `gorm:"column:product_ids"`
}
