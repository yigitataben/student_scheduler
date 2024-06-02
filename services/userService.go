package services

import (
	"github.com/yigitataben/student_scheduler/models"
	"github.com/yigitataben/student_scheduler/repositories"
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

func (us *UserService) GetUserByID(id int) (*models.User, error) {
	return us.UserRepository.GetUserByID(id)
}

func (us *UserService) UpdateUserByID(id int, email, password string) error {
	return us.UserRepository.UpdateUserByID(id, email, password)
}

func (us *UserService) DeleteUserByID(id int) error {
	return us.UserRepository.DeleteUserByID(id)
}
