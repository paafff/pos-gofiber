package repositories

import (
	"pos-gofiber/internal/domain/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// Fungsi NewUserRepository menerima parameter db yang merupakan pointer ke gorm.DB
// Buat objek UserRepository baru dengan db sebagai salah satu propertinya
// Kembalikan pointer ke objek UserRepository yang baru dibuat
func NewUserRepository(db *gorm.DB) *UserRepository {
	// Membuat dan mengembalikan objek UserRepository baru
	return &UserRepository{DB: db}
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser retrieves a user by ID from the database
func (r *UserRepository) GetUser(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user in the database
func (r *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser deletes a user by ID from the database
func (r *UserRepository) DeleteUser(id uint) error {
	if err := r.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetUsers retrieves all users from the database
func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
