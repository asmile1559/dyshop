package model

import (
	"time"
)

type User struct {
	UserID   int64  `gorm:"primaryKey;column:user_id" json:"user_id"`
	Email    string `gorm:"unique;column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Name     string `gorm:"column:name" json:"name"`
	Sign     string `gorm:"column:sign" json:"sign"`
	Url      string `gorm:"column:url" json:"url"` // 这里修改字段名
	Role     string `gorm:"column:role" json:"role"`
	Gender   string `gorm:"column:gender" json:"gender"`
	Birthday time.Time `gorm:"column:birthday" json:"birthday"`
	Phone    string `gorm:"column:phone" json:"phone"`
}
