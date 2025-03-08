package repository

import (
	"fmt"

	"github.com/ankush109/go-blog/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("user: %v\n", user.ID)
	return &user, nil
}

func NewRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
