package repositories

import (
	"pos-gofiber/internal/domain/models"

	"gorm.io/gorm"
)

type OrderedProductRepository struct {
	DB *gorm.DB
}

func NewOrderedProductRepository(db *gorm.DB) *OrderedProductRepository {
	return &OrderedProductRepository{DB: db}
}

// CreateOrderedProduct creates a new ordered product in the database
func (r *OrderedProductRepository) CreateOrderedProduct(orderedProduct *models.OrderedProduct) (*models.OrderedProduct, error) {
	if err := r.DB.Create(orderedProduct).Error; err != nil {
		return nil, err
	}
	return orderedProduct, nil
}

// GetOrderedProduct retrieves an ordered product by ID from the database
func (r *OrderedProductRepository) GetOrderedProduct(id uint) (*models.OrderedProduct, error) {
	var orderedProduct models.OrderedProduct
	if err := r.DB.First(&orderedProduct, id).Error; err != nil {
		return nil, err
	}
	return &orderedProduct, nil
}

// UpdateOrderedProduct updates an existing ordered product in the database
func (r *OrderedProductRepository) UpdateOrderedProduct(orderedProduct *models.OrderedProduct) (*models.OrderedProduct, error) {
	if err := r.DB.Save(orderedProduct).Error; err != nil {
		return nil, err
	}
	return orderedProduct, nil
}

// DeleteOrderedProduct deletes an ordered product by ID from the database
func (r *OrderedProductRepository) DeleteOrderedProduct(id uint) error {
	if err := r.DB.Delete(&models.OrderedProduct{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetOrderedProducts retrieves all ordered products from the database
func (r *OrderedProductRepository) GetOrderedProducts() ([]models.OrderedProduct, error) {
	var orderedProducts []models.OrderedProduct
	if err := r.DB.Find(&orderedProducts).Error; err != nil {
		return nil, err
	}
	return orderedProducts, nil
}
