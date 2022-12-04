package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DeletedAt gorm.DeletedAt
}
