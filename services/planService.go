package services

import (
	"errors"
	"gorm.io/gorm"

	"github.com/yigitataben/student_scheduler/models"
	"github.com/yigitataben/student_scheduler/repositories"
	"github.com/yigitataben/student_scheduler/requests"
)

var (
	ErrPlanNotFound = errors.New("plan not found")
)

type PlanService struct {
	PlanRepository *repositories.PlanRepository
}

func NewPlanService(planRepository *repositories.PlanRepository) *PlanService {
	return &PlanService{PlanRepository: planRepository}
}

func (s *PlanService) CreatePlan(planRequest requests.CreatePlanRequest) error {
	plan := models.Plan{
		LectureID: planRequest.LectureID,
		UserID:    planRequest.UserID,
		StartTime: planRequest.StartTime,
		EndTime:   planRequest.EndTime,
	}
	return s.PlanRepository.Create(&plan)
}

func (s *PlanService) GetAllPlans() ([]models.Plan, error) {
	return s.PlanRepository.FindAll()
}

func (s *PlanService) GetPlanByID(id string) (*models.Plan, error) {
	plan, err := s.PlanRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPlanNotFound
		}
		return nil, err
	}
	return plan, nil
}

func (s *PlanService) UpdatePlan(id string, newPlanData requests.CreatePlanRequest) (*models.Plan, error) {
	plan, err := s.PlanRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPlanNotFound
		}
		return nil, err
	}
	plan.LectureID = newPlanData.LectureID
	plan.UserID = newPlanData.UserID
	plan.StartTime = newPlanData.StartTime
	plan.EndTime = newPlanData.EndTime
	err = s.PlanRepository.Save(plan)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (s *PlanService) DeletePlan(id string) error {
	plan, err := s.PlanRepository.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPlanNotFound
		}
		return err
	}
	return s.PlanRepository.Delete(plan)
}
