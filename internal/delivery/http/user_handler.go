package http

import (
	"fmt"
	"net/http"

	"github.com/ankush109/go-blog/internal/domain"
	"github.com/ankush109/go-blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase usecase.UserUseCase
}

func NewUserHandler(r *gin.Engine, usecase usecase.UserUseCase) {
	userHandler := &UserHandler{usecase}
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
}

func (u *UserHandler) Register(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	err := u.usecase.Register(user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "User created successfully!"})

}

func (u *UserHandler) Login(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request!"})
		return
	}
	accessToken, err := u.usecase.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad"})
		return
	}
	// create a access Token :
	fmt.Println(accessToken, "access token")
	c.JSON(200, gin.H{"accessToken": accessToken})

}
