package models

import (
	"time"
)

type Transaction struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	PaymentMethod       string    `gorm:"size:255" json:"payment_method"`
	CustomerName        string    `gorm:"size:255" json:"customer_name"`
	Description         string    `gorm:"size:255" json:"description"`
	CreatedByID         string    `gorm:"size:255" json:"created_by_id"`
	ReceivedAmount      int       `json:"received_amount"`
	ReturnAmount        int       `json:"return_amount"`
	TotalAmount         int       `json:"total_amount"`
	TotalOrderedProduct int       `json:"total_ordered_product"`
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
