package database

import (
	"gorm.io/gorm"
	"jwt-go/internal/models"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&models.User{}, &models.Location{}, &models.Weather{})

	return nil
}

func Drop(db *gorm.DB) error {

	db.Migrator().DropTable("users")
	db.Migrator().DropTable("locations")
	db.Migrator().DropTable("weathers")

	return nil
}
