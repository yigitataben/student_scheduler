package services

import (
	"github.com/yigitataben/student_scheduler/models"
	"github.com/yigitataben/student_scheduler/repositories"
)

type PlanService struct {
	PlanRepository *repositories.PlanRepository
}

func NewPlanService(planRepository *repositories.PlanRepository) *PlanService {
	return &PlanService{PlanRepository: planRepository}
}

func (ps *PlanService) CreatePlan(LectureID int, UserID int, StartTime string, EndTime string, Status string) error {
	plan := models.Plan{
		LectureID: LectureID,
		UserID:    UserID,
		StartTime: StartTime,
		EndTime:   EndTime,
		Status:    Status,
	}
	return ps.PlanRepository.Create(&plan)
}

func (ps *PlanService) GetAllPlans() ([]models.Plan, error) {
	return ps.PlanRepository.GetAllPlans()
}

func (ps *PlanService) GetPlanByID(id int) (*models.Plan, error) {
	return ps.PlanRepository.GetPlanByID(id)
}

func (ps *PlanService) UpdatePlanByID(id int, lectureID int, userID int, startTime string, endTime string, status string) error {
	return ps.PlanRepository.UpdatePlanByID(id, lectureID, userID, startTime, endTime, status)
}

func (ps *PlanService) DeletePlanByID(id int) error {
	return ps.PlanRepository.DeletePlanByID(id)
}
