package service

import (
	"errors"
	"time"

	"github.com/okaraahmetoglu/go-clean-architecture/internal/app/auth/dto"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/okaraahmetoglu/go-clean-architecture/internal/domain/entity"
)

type AuthService interface {
	Authenticate(dto.LoginRequest) (*dto.LoginResponse, error)
	ValidateToken(tokenString string) (*entity.JWTClaims, error)
}

type authService struct {
	userRepo repository.UserRepository
	secret   string
}

func NewAuthService(userRepo repository.UserRepository, secret string) AuthService {
	return &authService{userRepo: userRepo, secret: secret}
}

func (u *authService) Authenticate(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.userRepo.GetUserByUsername(req.Username)
	if err != nil || user.Password != "hashedpassword" { // Normalde hash'lenmiş şifre kontrol edilmeli
		return nil, errors.New("invalid credentials")
	}

	claims := entity.JWTClaims{
		UserID: int64(user.ID),
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(u.secret))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: signedToken}, nil
}

func (u *authService) ValidateToken(tokenString string) (*entity.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &entity.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*entity.JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
