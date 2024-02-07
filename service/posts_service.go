package service

import (
	"go_gin/dto"
	"go_gin/entity"
	errorhandler "go_gin/errorHandler"
	"go_gin/repository"
)

type PostService interface {
	Create(req *dto.PostRequest) error
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService (r repository.PostRepository) *postService{
	return &postService{
		repository: r,
	}
}


func (s *postService) Create(req *dto.PostRequest) error {
	post := entity.Posts{
		UserID: req.UserID,
		Post: req.Post,
	}

	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.Create(&post); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}