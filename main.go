package main

import (
	"github.com/ankush109/go-blog/config"
	"github.com/ankush109/go-blog/internal/delivery/http"
	"github.com/ankush109/go-blog/internal/repository"
	"github.com/ankush109/go-blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB() // Connect to database
	r := gin.Default()       // Initialize Gin

	// Dependency injection
	userRepo := repository.NewRepository(db)
	userUseCase := usecase.NewUseCaseRepository(userRepo)
	http.NewUserHandler(r, userUseCase)

	postRepo := repository.NewPostRepository(db)
	postUseCase := usecase.NewPostUseCase(postRepo)
	http.NewPostHandler(r, postUseCase)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "server is rnning"})
	})

	// Start the server
	r.Run(":8080") // Add this line to run on port 8080
}
