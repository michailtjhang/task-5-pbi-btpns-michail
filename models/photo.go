package models

import "time"

type Photo struct {
	Title     string     `gorm:"not null"`
	Caption   string
	PhotoUrl  string     `gorm:"not null"`
	UserID    *uint64    `gorm:"foreignKey:UserID"`
	User      *User      `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
}