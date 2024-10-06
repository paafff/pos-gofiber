package services

import (
	"pos-gofiber/internal/domain/models"
	"pos-gofiber/internal/repositories"
)

type TransactionService struct {
	transactionRepository *repositories.TransactionRepository
}

func NewTransactionService(transactionRepository *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepository: transactionRepository}
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return s.transactionRepository.CreateTransaction(transaction)
}

func (s *TransactionService) GetTransaction(id uint) (*models.Transaction, error) {
	return s.transactionRepository.GetTransaction(id)
}

func (s *TransactionService) UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return s.transactionRepository.UpdateTransaction(transaction)
}

func (s *TransactionService) DeleteTransaction(id uint) error {
	return s.transactionRepository.DeleteTransaction(id)
}

func (s *TransactionService) GetTransactions() ([]models.Transaction, error) {
	return s.transactionRepository.GetTransactions()
}
