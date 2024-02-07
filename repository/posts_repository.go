package repository

import (
	"go_gin/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Posts) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(post *entity.Posts) error {
	err := r.db.Create(&post).Error
	return err
}