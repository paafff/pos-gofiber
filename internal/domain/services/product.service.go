package services

import (
	"pos-gofiber/internal/domain/models"
	"pos-gofiber/internal/repositories"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	return s.productRepository.CreateProduct(product)
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	return s.productRepository.GetProduct(id)
}

func (s *ProductService) UpdateProduct(product *models.Product) (*models.Product, error) {
	return s.productRepository.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.productRepository.DeleteProduct(id)
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.productRepository.GetProducts()
}
