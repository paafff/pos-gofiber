package services

import (
	"pos-gofiber/internal/domain/models"
	"pos-gofiber/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.userRepository.GetUser(id)
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return s.userRepository.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepository.DeleteUser(id)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.userRepository.GetUsers()
}
