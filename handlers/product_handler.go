package handlers

import (
	"go-photoUser/models"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var photos []models.Photo
		var wg sync.WaitGroup

		// Menambahkan satu goroutine ke WaitGroup
		wg.Add(1)

		// Memulai goroutine untuk melakukan operasi yang membutuhkan waktu lama
		go func() {
			defer wg.Done() // Menandai bahwa goroutine telah selesai
			db.Find(&photos)
		}()

		// Menunggu goroutine selesai
		wg.Wait()

		c.JSON(200, photos)
	}
}

func GetPhoto(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var photos models.Photo
		if err := db.First(&photos, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Photos not found"})
			return
		}
		c.JSON(200, photos)
	}
}

func CreatePhoto(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Photo
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Create(&input)
		c.JSON(201, input)
	}
}

func UpdatePhoto(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var photos models.Photo
		if err := db.First(&photos, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Photos not found"})
			return
		}

		var input models.Photo
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Model(&photos).Updates(input)
		c.JSON(200, photos)
	}
}

func DeletePhoto(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var photos models.Photo
		if err := db.First(&photos, id).Error; err != nil {
			c.JSON(404, gin.H{"message": "Photos not found"})
			return
		}

		db.Delete(&photos)
		c.JSON(200, gin.H{"message": "Photos deleted"})
	}
}