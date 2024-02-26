package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	PhotoID   *uint64   `gorm:"foreignKey:PhotoID"`
	Photo     *Photo    `gorm:"foreignKey:PhotoID;references:ID"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
}
