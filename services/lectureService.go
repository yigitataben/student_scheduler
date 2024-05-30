package services

import (
	"errors"

	"github.com/yigitataben/student_scheduler/models"
	"github.com/yigitataben/student_scheduler/repositories"
	"github.com/yigitataben/student_scheduler/requests"
)

var ErrLectureNotFound = errors.New("lecture not found")

type LectureService struct {
	LectureRepository *repositories.LectureRepository
}

func NewLectureService(lectureRepository *repositories.LectureRepository) *LectureService {
	return &LectureService{LectureRepository: lectureRepository}
}

func (s *LectureService) CreateLectures(lectureRequests []requests.CreateLectureRequest) error {
	var lectures []models.Lecture
	for _, input := range lectureRequests {
		lecture := models.Lecture{
			LectureName: input.LectureName,
		}
		lectures = append(lectures, lecture)
	}
	return s.LectureRepository.Create(lectures)
}

func (s *LectureService) GetAllLectures() ([]models.Lecture, error) {
	return s.LectureRepository.FindAll()
}

func (s *LectureService) GetLecture(id string) (*models.Lecture, error) {
	lecture, err := s.LectureRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, repositories.ErrLectureRecordNotFound) {
			return nil, ErrLectureNotFound
		}
		return nil, err
	}
	return lecture, nil
}

func (s *LectureService) DeleteLecture(id string) error {
	err := s.LectureRepository.Delete(id)
	if err != nil {
		if errors.Is(err, repositories.ErrLectureRecordNotFound) {
			return ErrLectureNotFound
		}
		return err
	}
	return nil
}
