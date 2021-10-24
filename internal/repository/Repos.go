package repository

import (
	"gorm.io/gorm"
	"jwt-go/internal/models"
)

type User interface {
	GetAll() ([]models.User, error)
	GetList(id int) (models.User, error)
	Delete(id int) (models.User, error)
	Update(id int, data models.User) (models.User, error)
}

type Auth interface {
	SignUp(data models.User) (models.User, error)
	GetUser(username, password string) (models.User, error)
}

type Repos struct {
	User
	Auth
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		User: NewUserRepo(db),
		Auth: NewAuthRepo(db),
	}
}
