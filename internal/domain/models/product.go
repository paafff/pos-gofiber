package models

import (
	"time"
)

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CreatedByID uint      `json:"created_by_id"`
	Name        string    `gorm:"size:255" json:"name"`
	ImageUrl    string    `gorm:"size:255" json:"image_url"`
	Stock       int       `json:"stock"`
	Sold        int       `json:"sold"`
	Price       float64   `json:"price"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
