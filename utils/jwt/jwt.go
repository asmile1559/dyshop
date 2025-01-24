package jwt

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	UserID  uint32
	Subject string
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint32) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("No Such Environment Variable")
	}

	jwtSecret := []byte(secret)
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	claims := User{
		userID,
		fmt.Sprintf("%08x", userID),
		jwt.RegisteredClaims{
			Issuer:    "dyshop-auth",
			ExpiresAt: jwt.NewNumericDate(expireTime),
			NotBefore: jwt.NewNumericDate(nowTime), // 在给定时间戳之前，JWT无效。
			IssuedAt:  jwt.NewNumericDate(nowTime),
			//Subject:   "", //  可以是一个用户ID、用户名或其他唯一标识符，目的是帮助识别JWT的持有者。
			//Audience:  nil, // 如果 aud 匹配：解码将成功，返回正确的 JWT payload, 否则则不会返回正确的 JWT payload
			//ID:        "", // 是一个唯一标识符，用于标识JWT令牌的唯一性。用于防止重放攻击、令牌撤销和令牌追踪
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*User, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &User{}, func(token *jwt.Token) (interface{}, error) {
		err := godotenv.Load()
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			logrus.Error("No Such Environment Variable")
			return nil, err
		}

		jwtSecret := []byte(secret)
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*User); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
