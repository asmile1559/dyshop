package service

import (
	"context"
	"errors"

	"github.com/asmile1559/dyshop/app/frontend/internal/model/user_model"
	"github.com/asmile1559/dyshop/app/frontend/internal/model/user_model/dto"

	"github.com/asmile1559/dyshop/utils/jwt"

	"github.com/asmile1559/dyshop/app/frontend/internal/dao/user_dao"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) Register(req *dto.RegisterRequest) (resp *dto.RegisterResponse, err error) {
	var count int64
	if err = user_dao.DB.Model(&user_model.User{}).Where("email = ?", req.Email).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &user_model.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     "user",
		Status:   1,
	}

	return nil, user_dao.DB.Create(user).Error
}

func (s *UserService) Login(req *dto.LoginRequest) (resp *dto.LoginResponse, err error) {
	var user user_model.User
	if err = user_dao.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := jwt.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token:  token,
		UserID: user.ID,
	}, nil
}

// optional
//func (s *UserService) Logout(req *dto.LogoutRequest) (resp dto.LogoutResponse, err error) {
//	return
//}
