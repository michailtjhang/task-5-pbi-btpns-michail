package seeders

import (
	"go-photoUser/models"

	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) {
	// Seeder untuk Users
	user1 := models.User{Username: "user1", Email: "user1@example.com", Password: "password1"}
	user2 := models.User{Username: "user2", Email: "user2@example.com", Password: "password2"}

	db.Create(&user1)
	db.Create(&user2)

	// Seeder untuk Photos
	photo1 := models.Photo{Title: "Photo 1", PhotoUrl: "https://example.com/photo1.jpg"}
	photo2 := models.Photo{Title: "Photo 2", PhotoUrl: "https://example.com/photo2.jpg"}

	db.Create(&photo1)
	db.Create(&photo2)
}
