package dto

type RegisterRequest struct {
	Username        string `json:"username" form:"username" binding:"required,min=3,max=50"`
	Password        string `json:"password" form:"password" binding:"required,min=6,max=50"`
	Email           string `json:"email" form:"email" binding:"required,email"`
	Phone           string `json:"phone" form:"phone" binding:"omitempty,len=11"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,eqfield=Password"`
}

type RegisterResponse struct {
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=50"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserID int64  `json:"userId"`
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Role      string `json:"role"`
}
