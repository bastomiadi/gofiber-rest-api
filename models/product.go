package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string `'json:"name"`
	SerialNumber string `json:"serial_number"`
	DeletedAt    gorm.DeletedAt
}
