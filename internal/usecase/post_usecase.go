package usecase

import (
	"github.com/ankush109/go-blog/internal/domain"
	"github.com/ankush109/go-blog/internal/repository"
)

type PostUseCase interface {
	CreatePost(title string, content string, userId uint) error
	GetPostById(id uint) (*domain.Post, error)
	GetPostsByUserId(userId uint) ([]domain.Post, error)
	DeletePostById(id uint) error
	UpdatePostById(id uint, title string, content string) error
}

type postUseCase struct {
	repo repository.PostRepository
}

// CreatePost implements PostUseCase.
func (p *postUseCase) CreatePost(title string, content string, userId uint) error {
	post := &domain.Post{
		Title:   title,
		Content: content,
		UserID:  userId,
	}
	return p.repo.CreatePost(post)
}

// DeletePostById implements PostUseCase.
func (p *postUseCase) DeletePostById(id uint) error {
	return p.repo.DeletePostById(id)
}

// GetPostById implements PostUseCase.
func (p *postUseCase) GetPostById(id uint) (*domain.Post, error) {
	return p.repo.GetPostById(id)
}

// GetPostsByUserId implements PostUseCase.
func (p *postUseCase) GetPostsByUserId(userId uint) ([]domain.Post, error) {
	return p.repo.GetPostsByUserId(userId)
}

// UpdatePostById implements PostUseCase.
func (p *postUseCase) UpdatePostById(id uint, title string, content string) error {
	return p.repo.UpdatePostById(id, &domain.Post{Title: title, Content: content})
}

func NewPostUseCase(repo repository.PostRepository) PostUseCase {
	return &postUseCase{repo}
}
