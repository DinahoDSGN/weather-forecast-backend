package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"jwt-go/internal/models"
	"jwt-go/internal/repository"
	"time"
)

const salt = "jglkdabgfa987r89sahdnlkn"
const signingKey = "gpldmzlkfgh87809"

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(data models.User) (models.User, error) {
	data.Password = generatePasswordHash(data.Password)
	return s.repo.SignUp(data)
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	data, err := s.repo.GetUser(username, generatePasswordHash(password))
	if data.Id == 0 {
		return "Invalid username or password", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		data.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return []byte(signingKey), nil
		})
	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *TokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
