package services

import (
	"pos-gofiber/internal/domain/models"
	"pos-gofiber/internal/repositories"
)

type OrderedProductService struct {
	orderedProductRepository *repositories.OrderedProductRepository
}

func NewOrderedProductService(orderedProductRepository *repositories.OrderedProductRepository) *OrderedProductService {
	return &OrderedProductService{orderedProductRepository: orderedProductRepository}
}

func (s *OrderedProductService) CreateOrderedProduct(orderedProduct *models.OrderedProduct) (*models.OrderedProduct, error) {
	return s.orderedProductRepository.CreateOrderedProduct(orderedProduct)
}

func (s *OrderedProductService) GetOrderedProduct(id uint) (*models.OrderedProduct, error) {
	return s.orderedProductRepository.GetOrderedProduct(id)
}

func (s *OrderedProductService) UpdateOrderedProduct(orderedProduct *models.OrderedProduct) (*models.OrderedProduct, error) {
	return s.orderedProductRepository.UpdateOrderedProduct(orderedProduct)
}

func (s *OrderedProductService) DeleteOrderedProduct(id uint) error {
	return s.orderedProductRepository.DeleteOrderedProduct(id)
}

func (s *OrderedProductService) GetOrderedProducts() ([]models.OrderedProduct, error) {
	return s.orderedProductRepository.GetOrderedProducts()
}
