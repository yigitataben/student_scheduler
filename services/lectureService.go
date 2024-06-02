package services

import (
	"github.com/yigitataben/student_scheduler/models"
	"github.com/yigitataben/student_scheduler/repositories"
	"github.com/yigitataben/student_scheduler/requests"
)

type LectureService struct {
	LectureRepository *repositories.LectureRepository
}

func NewLectureService(lectureRepository *repositories.LectureRepository) *LectureService {
	return &LectureService{LectureRepository: lectureRepository}
}

func (ls *LectureService) CreateLectures(lectureRequests []requests.CreateLectureRequest) error {
	var lectures []models.Lecture
	for _, request := range lectureRequests {
		lecture := models.Lecture{
			LectureName: request.LectureName,
		}
		lectures = append(lectures, lecture)
	}
	return ls.LectureRepository.Create(lectures)
}

func (ls *LectureService) GetAllLectures() ([]models.Lecture, error) {
	return ls.LectureRepository.GetAllLectures()
}

func (ls *LectureService) GetLectureByID(id int) (*models.Lecture, error) {
	return ls.LectureRepository.GetLectureByID(id)
}

func (ls *LectureService) UpdateLectureByID(id int, lectureName string) error {
	return ls.LectureRepository.UpdateLectureByID(id, lectureName)
}

func (ls *LectureService) DeleteLectureByID(id int) error {
	return ls.LectureRepository.DeleteLectureByID(id)
}
