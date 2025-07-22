package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"task4/internal/config"
	"task4/internal/model"
	"time"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(user model.User) (string, error) {
	tokenExpiry := config.GetConfig().Auth.TokenExpiry
	expirationTime := time.Now().Add(time.Duration(tokenExpiry) * time.Second)
	claims := &Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetConfig().Auth.JwtSecret))
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().Auth.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
