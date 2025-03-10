package models

import "time"

type User struct {
  ID        int    `gorm:"primaryKey"`
  Name      string `gorm:"varchar(100)"`
  Email     string `gorm:"varchar(100)"`
  Password  string `gorm:"varchar(255)"`
  Gender    string `gorm:"varchar(10)"`
  CreatedAt time.Time
  UpdatedAt time.Time
}
