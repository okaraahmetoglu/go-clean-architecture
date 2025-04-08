package entity

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
