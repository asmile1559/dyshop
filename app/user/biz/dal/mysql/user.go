package mysql

import (
	"github.com/asmile1559/dyshop/app/user/biz/model"
)

// GetUserByEmail 根据邮箱查询用户
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser 存储用户信息到数据库
func CreateUser(user *model.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}