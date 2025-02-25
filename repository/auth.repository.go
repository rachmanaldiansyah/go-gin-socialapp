package repository

import (
	"gorm.io/gorm"

	"go-gin-sosmed/models"
)

type AuthRepository interface{
  EmailExist(email string) bool
  Register(req *models.User) error
}

type authRepository struct {
  db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
  return &authRepository{
    db: db,
  }
}

func (r *authRepository) Register(user *models.User) error {
  err := r.db.Create(&user).Error

  return err
}

func (r *authRepository) EmailExist(email string) bool {
  var user models.User
  err := r.db.First(&user, "email = ?", email).Error

  return err == nil
}
