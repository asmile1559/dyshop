package mysql

import (
	"fmt"

	"github.com/asmile1559/dyshop/app/user/biz/model"
	"gorm.io/gorm"
)

// GetUserByEmail 根据邮箱查询用户
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID 根据ID查询用户
func GetUserByID(userID uint32) (*model.User, error) {
	var user model.User
	err := db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
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

// UpdateUser 更新用户信息
func UpdateUser(user *model.User) error {
	return db.Save(user).Error
}

// DeleteUserByID 根据用户ID删除用户
func DeleteUserByID(userID uint32) error {
	var user model.User
	// 删除用户
	if err := db.Where("user_id = ?", userID).Delete(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("用户不存在")
		}
		return err
	}
	return nil
}