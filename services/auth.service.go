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
  Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
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

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
  var data dto.LoginResponse

  user, err := s.repository.GetUserByEmail(req.Email)
  if err != nil {
    return nil, &exceptions.NotFoundError{Message: "wrong email or password"}
  }

  if err := helpers.VerifyPassword(user.Password, req.Password); err != nil {
    return nil, &exceptions.NotFoundError{Message: "worng email or password"}
  }

  token, err := helpers.GenerateToken(user)
  if err != nil {
    return nil, &exceptions.InternalServerError{Message: err.Error()}
  }

  data = dto.LoginResponse{
    ID:     user.ID,
    Name:   user.Name,
    Token:  token,
  }

  return &data, nil
}
