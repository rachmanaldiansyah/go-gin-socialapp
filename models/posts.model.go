package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
  ID          int       `gorm:"primaryKey"`
  UserID      int
  User        User      `gorm:"foreignKey:UserID"`
  Tweet       string    `gorm:"text"`
  PictureUrl  *string   `gorm:"text"`
  CreatedAt   time.Time
  UpdatedAt   time.Time
}

type PostRepository interface {
  Create(post *Post) error
}

type postRepository struct {
  db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
  return &postRepository{
    db: db,
  }
}

func (r *postRepository) Create(post *Post) error {
  err := r.db.Create(&post).Error
  return err
}