package model

import (
	"time"

)

type User struct {
	
	UserID    uint      `gorm:"primarykey"`  // 自定义主键，使用 user_id
	Email    string `gorm:"unique;not null"` // 唯一且非空的邮箱
	Password string `gorm:"not null"`        // 密码字段
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"` // 软删除
}
