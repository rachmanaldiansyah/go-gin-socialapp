package services

import (
	"go-gin-sosmed/dto"
	"go-gin-sosmed/exceptions"
	"go-gin-sosmed/helpers"
	"go-gin-sosmed/models"
	"go-gin-sosmed/repository"
)

type AuthService interface{
  Register(req *dto.RegisterRequest) error
}

type authService struct {
  repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
  return &authService{
    repository: r,
  }
}

func (s *authService) Register(req *dto.RegisterRequest) error {
  if emailExist := s.repository.EmailExist(req.Email); emailExist {
    return &exceptions.BadRequestError{Message: "email already registered"}
  }

  if req.Password != req.PasswordConfirmation {
    return &exceptions.BadRequestError{Message: "password not match"}
  }

  passwordHash, err := helpers.HashPassword(req.Password)
  if err != nil {
    return &exceptions.InternalServerError{Message: err.Error()}
  }

  user := models.User{
    Name: req.Name,
    Email: req.Email,
    Password: passwordHash,
    Gender: req.Gender,
  }

  if err := s.repository.Register(&user); err != nil {
    return &exceptions.InternalServerError{Message: err.Error()}
  }

  return nil
}
