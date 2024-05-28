package services

import (
	"errors"
	"github.com/yigitataben/student_scheduler/models"
	"github.com/yigitataben/student_scheduler/repositories"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepo}
}

func (s *UserService) SignUp(email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{Email: email, Password: string(hash)}
	return s.UserRepository.Create(&user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepository.FindAll()
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	user, err := s.UserRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, repositories.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(id, email, password string) (*models.User, error) {
	user, err := s.UserRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, repositories.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Email = email
	user.Password = string(hash)
	if err := s.UserRepository.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(id string) error {
	err := s.UserRepository.Delete(id)
	if err != nil {
		if errors.Is(err, repositories.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}
