package usecase

import (
	"fmt"
	"time"

	"github.com/ankush109/go-blog/internal/domain"
	"github.com/ankush109/go-blog/internal/repository"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("ankush2003")

type UserUseCase interface {
	Register(email string, password string) error
	Login(email string, password string) (string, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

// Login implements UserUseCase.
func (u *userUseCase) Login(email string, password string) (string, error) {
	fmt.Println("in login usecase")
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found!")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Invalid credentials!")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	accessToken, err := token.SignedString(jwtSecret)
	return accessToken, nil
}

// Register implements UserUseCase.
func (u *userUseCase) Register(email string, password string) error {
	fmt.Println("in register usecase")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//
	fmt.Println(hashedPassword)
	user := &domain.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     "Ankush",
	}
	return u.repo.CreateUser(user)
}

func NewUseCaseRepository(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo}
}
