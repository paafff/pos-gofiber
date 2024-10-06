package models

import (
	"time"
)

type OrderedProduct struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TransactionID uint      `json:"transaction_id"`
	ProductID     uint      `json:"product_id"`
	Quantity      int       `json:"quantity"`
	Stock         int       `json:"stock"`
	Name          string    `gorm:"size:255" json:"name"`
	ImageUrl      string    `gorm:"size:255" json:"image_url"`
	Price         float64   `json:"price"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
