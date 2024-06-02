package repositories

import (
	"github.com/yigitataben/student_scheduler/models"
	"gorm.io/gorm"
)

type LectureRepository struct {
	DB *gorm.DB
}

func NewLectureRepository(db *gorm.DB) *LectureRepository {
	return &LectureRepository{DB: db}
}

func (lr *LectureRepository) Create(lectures []models.Lecture) error {
	return lr.DB.Create(&lectures).Error
}

func (lr *LectureRepository) GetAllLectures() ([]models.Lecture, error) {
	var users []models.Lecture
	err := lr.DB.Order("created_at desc").Find(&users).Error
	return users, err
}

func (lr *LectureRepository) GetLectureByID(id int) (*models.Lecture, error) {
	lecture := &models.Lecture{}
	result := lr.DB.First(lecture, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return lecture, nil
}

func (lr *LectureRepository) UpdateLectureByID(id int, lectureName string) error {
	lecture := &models.Lecture{}
	result := lr.DB.First(lecture, id)
	if result.Error != nil {
		return result.Error
	}
	lecture.LectureName = lectureName
	result = lr.DB.Save(lecture)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lr *LectureRepository) DeleteLectureByID(id int) error {
	result := lr.DB.Unscoped().Delete(&models.Lecture{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
