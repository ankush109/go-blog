package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ankush109/go-blog/internal/domain"
	"github.com/ankush109/go-blog/internal/middleware"
	"github.com/ankush109/go-blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	usecase usecase.PostUseCase
}

func NewPostHandler(r *gin.Engine, usecase usecase.PostUseCase) *PostHandler {
	postHandler := &PostHandler{usecase}
	authRoutes := r.Group("/post")
	authRoutes.Use(middleware.AuthMiddleware())
	authRoutes.POST("", postHandler.CreatePost)
	authRoutes.GET("", postHandler.GetPostsByUserId)
	authRoutes.GET("/:id", postHandler.GetPostById)
	authRoutes.PUT("/:id", postHandler.UpdatePostById)
	authRoutes.DELETE("/:id", postHandler.DeletePostById)
	return postHandler
}

func (p *PostHandler) GetPostsByUserId(c *gin.Context) {
	userId, _ := c.Get("user_id")
	convertedUserId := uint(userId.(float64))
	posts, err := p.usecase.GetPostsByUserId(convertedUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
func (p *PostHandler) UpdatePostById(c *gin.Context) {
	var post domain.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	postId := c.Param("id")
	convertedPostId, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	err = p.usecase.UpdatePostById(uint(convertedPostId), post.Title, post.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Post updated successfully!"})
}
func (p *PostHandler) CreatePost(c *gin.Context) {

	var post domain.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	userId, _ := c.Get("user_id")
	fmt.Println(userId, "user_id")
	convertedUserId := uint(userId.(float64))
	post.UserID = convertedUserId
	err := p.usecase.CreatePost(post.Title, post.Content, post.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Post created successfully!"})

}
func (p *PostHandler) DeletePostById(c *gin.Context) {
	id := c.Param("id")
	convertedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	err = p.usecase.DeletePostById(uint(convertedId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Post deleted successfully!"})
}
func (p *PostHandler) GetPostById(c *gin.Context) {
	id := c.Param("id")
	convertedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	post, err := p.usecase.GetPostById(uint(convertedId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}
