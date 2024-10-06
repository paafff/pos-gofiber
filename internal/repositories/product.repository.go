package repositories

import (
	"pos-gofiber/internal/domain/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// CreateProduct creates a new product in the database
func (r *ProductRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	if err := r.DB.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// GetProduct retrieves a product by ID from the database
func (r *ProductRepository) GetProduct(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates an existing product in the database
func (r *ProductRepository) UpdateProduct(product *models.Product) (*models.Product, error) {
	if err := r.DB.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// DeleteProduct deletes a product by ID from the database
func (r *ProductRepository) DeleteProduct(id uint) error {
	if err := r.DB.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetProducts retrieves all products from the database
func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
