package repositories

import (
	"pos-gofiber/internal/domain/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	if err := r.DB.Create(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) GetTransaction(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.DB.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	if err := r.DB.Save(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) DeleteTransaction(id uint) error {
	if err := r.DB.Delete(&models.Transaction{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepository) GetTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.DB.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) GetTransactionByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) GetTransactionByOrderID(orderID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.DB.Where("order_id = ?", orderID).First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) GetTransactionByOrderIDAndUserID(orderID, userID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.DB.Where("order_id = ? AND user_id = ?", orderID, userID).First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) GetTransactionByOrderIDAndUserIDAndStatus(orderID, userID uint, status string) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.DB.Where("order_id = ? AND user_id = ? AND status = ?", orderID, userID, status).First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) GetTransactionByOrderIDAndStatus(orderID uint, status string) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.DB.Where("order_id = ? AND status = ?", orderID, status).First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}
