package repositories

import (
	"github.com/yigitataben/student_scheduler/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) SignUp(user *models.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := ur.DB.Order("created_at desc").Find(&users).Error
	return users, err
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	result := ur.DB.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (ur *UserRepository) UpdateUserByID(id int, email, password string) error {
	user := &models.User{}
	result := ur.DB.First(user, id)
	if result.Error != nil {
		return result.Error
	}
	user.Email = email
	user.Password = password
	result = ur.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *UserRepository) DeleteUserByID(id int) error {
	result := ur.DB.Unscoped().Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
