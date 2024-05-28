package services

import (
	"errors"
	"github.com/yigitataben/student_scheduler/models"
	"github.com/yigitataben/student_scheduler/repositories"
)

var (
	ErrSchedulingConflict = errors.New("scheduling conflict")
	ErrPlanNotFound       = errors.New("plan not found")
)

type PlanService struct {
	PlanRepository repositories.PlanRepository
}

func NewPlanService(planRepository repositories.PlanRepository) *PlanService {
	return &PlanService{PlanRepository: planRepository}
}

func (s *PlanService) CreatePlan(plan *models.Plan) error {
	existingPlans, err := s.PlanRepository.FindConflictingPlans(plan.UserID, plan.StartTime, plan.EndTime)
	if err != nil {
		return err
	}

	if len(existingPlans) > 0 {
		return ErrSchedulingConflict
	}

	return s.PlanRepository.Create(plan)
}

func (s *PlanService) UpdatePlanStatus(id, status string) (*models.Plan, error) {
	plan, err := s.PlanRepository.FindByID(id)
	if err != nil {
		return nil, ErrPlanNotFound
	}

	plan.Status = status
	err = s.PlanRepository.Save(plan)
	return plan, err
}

func (s *PlanService) GetAllPlans() ([]models.Plan, error) {
	return s.PlanRepository.FindAll()
}

func (s *PlanService) GetPlanByID(id string) (*models.Plan, error) {
	return s.PlanRepository.FindByID(id)
}

func (s *PlanService) UpdatePlan(id string, newPlanData *models.Plan) (*models.Plan, error) {
	plan, err := s.PlanRepository.FindByID(id)
	if err != nil {
		return nil, ErrPlanNotFound
	}

	plan.LectureName = newPlanData.LectureName
	plan.UserID = newPlanData.UserID
	err = s.PlanRepository.Save(plan)
	return plan, err
}

func (s *PlanService) DeletePlan(id string) error {
	plan, err := s.PlanRepository.FindByID(id)
	if err != nil {
		return ErrPlanNotFound
	}

	return s.PlanRepository.Delete(plan)
}
