package domain

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User `gorm:"constraint:OnDelete:CASCADE;"`
}
