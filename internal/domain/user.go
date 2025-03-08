package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint `gorm:"primaryKey"`
	Email    string
	Name     string
	Password string
	Posts    []Post `gorm:"foreignKey:UserID"`
}
