package repository

import (
	"gorm.io/gorm"
	"jwt-go/internal/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (database *UserRepo) GetAll() ([]models.User, error) {
	var user []models.User

	database.db.Preload("Location").Model(&models.User{}).Scan(&user)

	return user, nil
}

func (database *UserRepo) GetList(id int) (models.User, error) {
	var user models.User

	database.db.Preload("Location").Raw("SELECT * FROM users WHERE id = ?", id).Find(&user)

	return user, nil
}

func (database *UserRepo) Delete(id int) (models.User, error) {
	user, _ := database.GetList(id)
	if user.Id == 0 {
		return user, nil
	}

	database.db.Delete(&user)

	return user, nil
}

func (database *UserRepo) Update(id int, data models.User) (models.User, error) {
	user := models.User{
		Username:  data.Username,
		Password:  data.Password,
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Location: &models.Location{
			City:   data.Location.City,
			Region: data.Location.Region,
		},
		Age: data.Age,
	}

	database.db.Model(&user).Where("id = ?", id).Updates(&user).Find(&user)

	if user.Id == 0 {
		return user, nil
	}

	database.db.Save(&user)

	return user, nil
}
