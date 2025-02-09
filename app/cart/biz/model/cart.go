package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint64 `gorm:"primaryKey"` // 自增主键
	UserId    uint64 `gorm:"index"`      // 用户ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`             // 支持软删除
	CartItems []CartItem     `gorm:"foreignKey:CartId"` // 一对多关系：Cart.id -> CartItem.cart_id
}

type CartItem struct {
	ID        uint64 `gorm:"primaryKey"` // 自增主键
	CartId    uint64 `gorm:"index"`      // 外键，对应 Cart.ID
	ProductId uint64 // 商品ID
	Quantity  int    // 数量
}
