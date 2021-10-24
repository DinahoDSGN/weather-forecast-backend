package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"jwt-go/internal/models"
	"jwt-go/internal/repository"
	"jwt-go/util"
	"log"
	"time"
)

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
	config, err := util.LoadConfig(".") // initialize config
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

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

	return token.SignedString([]byte(config.SIGNING_KEY))
}

func (s *AuthService) ParseToken(accessToken string) (uint, error) {
	config, err := util.LoadConfig(".") // initialize config
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return []byte(config.SIGNING_KEY), nil
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
	config, err := util.LoadConfig(".") // initialize config
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(config.SALT)))
}
