package repositories

import (
	"errors"
	"github.com/yigitataben/student_scheduler/models"
	"gorm.io/gorm"
)

var (
	ErrUserRecordNotFound = errors.New("record not found")
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Order("created_at desc").Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserRecordNotFound
	}
	return &user, err
}

func (r *UserRepository) Save(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(id string) error {
	var user models.User
	err := r.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrUserRecordNotFound
	}
	return r.DB.Unscoped().Delete(&user).Error
}
