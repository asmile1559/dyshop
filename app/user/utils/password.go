package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 加密用户密码
func HashPassword(password string) (string, error) {
	// 使用 bcrypt 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword 验证用户输入的密码是否与存储的哈希密码匹配
func VerifyPassword(storedPassword, enteredPassword string) bool {
	// 使用 bcrypt.CompareHashAndPassword 比较密码
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(enteredPassword))
	if err != nil {
		// 如果密码不匹配，返回 false
		return false
	}
	// 密码匹配，返回 true
	return true
}
