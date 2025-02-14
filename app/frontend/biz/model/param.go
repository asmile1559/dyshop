package model

//定义请求的参数结构体，用于参数校验
// ParamRegister 注册请求参数
type ParamRegister struct {
	Email	string	`json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// ParamLogout 登出请求参数
type ParamLogout struct {
	UserID	int64 `json:"user_id" binding:"required"`
}

// ParamUpdateUser 更新用户信息请求参数
type ParamUpdateUser struct {
	UserID	int64 `json:"user_id" binding:"required"`
	Email	string	`json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

// ParamDeleteUser 删除用户请求参数
type ParamDeleteUser struct {
	UserID	int64 `json:"user_id" binding:"required"`
}

