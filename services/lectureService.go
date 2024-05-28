package services

import (
	"errors"
	"github.com/yigitataben/student_scheduler/models"
	"github.com/yigitataben/student_scheduler/repositories"
)

var (
	ErrLectureNotFound = errors.New("lecture not found")
)

type LectureService struct {
	LectureRepository repositories.LectureRepository
}

func NewLectureService(lectureRepo repositories.LectureRepository) *LectureService {
	return &LectureService{LectureRepository: lectureRepo}
}

func (s *LectureService) CreateLectures(lectures []models.Lecture) error {
	return s.LectureRepository.Create(lectures)
}

func (s *LectureService) GetAllLectures() ([]models.Lecture, error) {
	return s.LectureRepository.FindAll()
}

func (s *LectureService) GetLecture(id string) (*models.Lecture, error) {
	lecture, err := s.LectureRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, repositories.ErrRecordNotFound) {
			return nil, ErrLectureNotFound
		}
		return nil, err
	}
	return lecture, nil
}

func (s *LectureService) DeleteLecture(id string) error {
	err := s.LectureRepository.Delete(id)
	if err != nil {
		if errors.Is(err, repositories.ErrRecordNotFound) {
			return ErrLectureNotFound
		}
		return err
	}
	return nil
}
