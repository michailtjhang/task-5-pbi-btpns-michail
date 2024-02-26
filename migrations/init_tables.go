package migrations

import (
	"go-photoUser/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Photo{},
		&models.User{},
	)
}