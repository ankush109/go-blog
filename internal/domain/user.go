package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string `gorm:"uniqueIndex;not null"`
	Name     string
	Password string
	Posts    []Post `gorm:"foreignKey:UserID"`
}
