package handlers

import (
	"go-photoUser/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UpdateUsername(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
  
	  // Dapatkan user ID dari context 
	  userId := c.MustGet("userId").(uint)
  
	  // Baca username baru dari request body
	  var input struct {
		Username string `json:"username"`
	  }
	  if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"message": "Invalid input"})
		return
	  }
  
	  // Cek apakah username sudah digunakan
	  var existingUser models.User
	  if err := db.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{"message": "Username already exists"})
		return
	  }
  
	  // Update username user
	  var user models.User
	  if err := db.Model(&user).Where("id = ?", userId).Update("username", input.Username).Error; err != nil {
		c.JSON(500, gin.H{"message": "Failed to update username"})
		return
	  }
  
	  c.JSON(200, gin.H{"message": "Username updated successfully"})
	}
  }

  func DeleteUser(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
	  
	  // Ambil userId dari URL parameter
	  userId := c.Param("userId")
  
	  // Hapus user berdasarkan userId
	  if err := db.Where("id = ?", userId).Delete(&models.User{}).Error; err != nil {
		c.JSON(500, gin.H{"message": "Failed to delete user"})
		return
	  }
  
	  c.JSON(200, gin.H{"message": "User deleted successfully"})
	}
  }