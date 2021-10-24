package service

import (
	"jwt-go/internal/models"
	"jwt-go/internal/repository"
)

type User interface {
	GetAll() ([]models.User, error)
	GetList(id int) (models.User, error)
	Delete(id int) (models.User, error)
	Update(id int, data models.User) (models.User, error)
}

type Auth interface {
	SignUp(data models.User) (models.User, error)
	ParseToken(accessToken string) (uint, error)
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	User
	Auth
}

func NewService(repos *repository.Repos) *Service {
	return &Service{
		User: NewUserService(repos.User),
		Auth: NewAuthService(repos.Auth),
	}
}
