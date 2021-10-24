package repository

import (
	"gorm.io/gorm"
	"jwt-go/internal/models"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (database *AuthRepo) SignUp(data models.User) (models.User, error) {
	user := models.User{
		Username: data.Username,
		Password: data.Password,
	}

	database.db.Create(&user)

	return user, nil
}

func (database *AuthRepo) GetUser(username, password string) (models.User, error) {
	var data models.User

	err := database.db.Raw("SELECT * FROM users WHERE username = ? AND password = ?", username, password).Find(&data)

	if data.Id == 0 {
		return data, err.Error
	}

	return data, nil
}
