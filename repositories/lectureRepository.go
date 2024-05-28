package repositories

import (
	"errors"
	"github.com/yigitataben/student_scheduler/models"
	"gorm.io/gorm"
)

var (
	ErrLectureRecordNotFound = errors.New("record not found")
)

type LectureRepository struct {
	DB *gorm.DB
}

func NewLectureRepository(db *gorm.DB) *LectureRepository {
	return &LectureRepository{DB: db}
}

func (r *LectureRepository) Create(lectures []models.Lecture) error {
	return r.DB.Create(&lectures).Error
}

func (r *LectureRepository) FindAll() ([]models.Lecture, error) {
	var lectures []models.Lecture
	err := r.DB.Order("id asc").Find(&lectures).Error
	return lectures, err
}

func (r *LectureRepository) FindByID(id string) (*models.Lecture, error) {
	var lecture models.Lecture
	err := r.DB.First(&lecture, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrLectureRecordNotFound
	}
	return &lecture, err
}

func (r *LectureRepository) Delete(id string) error {
	var lecture models.Lecture
	err := r.DB.First(&lecture, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrLectureRecordNotFound
	}
	return r.DB.Unscoped().Delete(&lecture).Error
}
