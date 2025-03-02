package services

import (
	"go-gin-sosmed/dto"
	"go-gin-sosmed/exceptions"
	"go-gin-sosmed/models"
)

type PostService interface{
	Create(req *dto.PostRequest) error
}

type postService struct {
	repository models.PostRepository
}

func NewPostService(r models.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) Create(req *dto.PostRequest) error {
	post := models.Post{
		UserID: req.UserID,
		Tweet: req.Tweet,
	}

	if req.PictureUrl != nil {
		post.PictureUrl = &req.PictureUrl.Filename
	}

	if err := s.repository.Create(&post); err != nil {
		return &exceptions.InternalServerError{Message: err.Error()}
	}

	return nil
}