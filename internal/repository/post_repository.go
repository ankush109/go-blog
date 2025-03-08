package repository

import (
	"github.com/ankush109/go-blog/internal/domain"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *domain.Post) error
	GetPostById(id uint) (*domain.Post, error)
	GetPostsByUserId(userId uint) ([]domain.Post, error)
	DeletePostById(id uint) error
	UpdatePostById(id uint, post *domain.Post) error
}

type postRepository struct {
	db *gorm.DB
}

// CreatePost implements PostRepository.
func (p *postRepository) CreatePost(post *domain.Post) error {
	return p.db.Create(post).Error
}

// GetPostById implements PostRepository.
func (p *postRepository) GetPostById(id uint) (*domain.Post, error) {
	var post domain.Post
	err := p.db.Select("id, title, content, user_id, created_at, updated_at").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// GetPostsByUserId implements PostRepository.
func (p *postRepository) GetPostsByUserId(userId uint) ([]domain.Post, error) {
	var posts []domain.Post
	err := p.db.Preload("User").Where("user_id = ?", userId).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// UpdatePostById implements PostRepository.
func (p *postRepository) UpdatePostById(id uint, post *domain.Post) error {

	return p.db.Model(&domain.Post{}).Where("id = ?", id).Updates(post).Error

}
func (p *postRepository) DeletePostById(id uint) error {
	return p.db.Delete(&domain.Post{}, id).Error
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}
