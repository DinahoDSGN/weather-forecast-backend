package database

import (
	"gorm.io/gorm"
	"jwt-go/internal/models"
)

func Migrate(db *gorm.DB) error {

	db.Migrator().DropTable("messages")
	db.Migrator().DropTable("chats")
	db.AutoMigrate(&models.User{}, &models.Location{})

	return nil
}

func Drop(db *gorm.DB) error {

	err := db.Migrator().DropTable("users")
	if err != nil {
		return err
	}

	return nil
}
