package services

import (
	"errors"
	"github.com/yigitataben/student_scheduler/models"
	"github.com/yigitataben/student_scheduler/repositories"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (us *UserService) SignUp(email, password string) error {
	user := models.User{
		Email:    email,
		Password: password,
	}
	return us.UserRepository.SignUp(&user)
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	return us.UserRepository.GetAllUsers()
}

func (us *UserService) GetUserByID(id uint) (*models.User, error) {
	return us.UserRepository.GetUserByID(id)
}

func (us *UserService) UpdateUserByID(id uint, email, password string) error {
	return us.UserRepository.UpdateUserByID(id, email, password)
}

func (us *UserService) DeleteUserByID(id uint) error {
	return us.UserRepository.DeleteUserByID(id)
}
