package model

// AuthBase 复用 Email 和 Password 字段
type AuthBase struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

// ParamRegister 注册请求参数
type ParamRegister struct {
	AuthBase
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	AuthBase
}