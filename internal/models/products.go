package models

import (
	"time"
)

// Product model
type Product struct {
	ID        uint      `gorm:"primary_key;column:id" json:"id"`
	Name      string    `gorm:"column:name;" json:"name"`
	Code      string    `gorm:"column:code;" json:"code"`
	Price     float64   `gorm:"column:price" json:"price"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
