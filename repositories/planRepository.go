package repositories

import (
	"github.com/yigitataben/student_scheduler/models"
	"gorm.io/gorm"
)

type PlanRepository struct {
	DB *gorm.DB
}

func NewPlanRepository(db *gorm.DB) *PlanRepository {
	return &PlanRepository{DB: db}
}

func (pr *PlanRepository) Create(plan *models.Plan) error {
	return pr.DB.Create(plan).Error
}

func (pr *PlanRepository) GetAllPlans() ([]models.Plan, error) {
	var plans []models.Plan
	err := pr.DB.Order("created_at desc").Find(&plans).Error
	return plans, err
}

func (pr *PlanRepository) GetPlanByID(id int) (*models.Plan, error) {
	plan := &models.Plan{}
	result := pr.DB.First(plan, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return plan, nil
}

func (pr *PlanRepository) UpdatePlanByID(id int, lectureID int, userID int, startTime string, endTime string, status string) error {
	plan := &models.Plan{}
	result := pr.DB.First(plan, id)
	if result.Error != nil {
		return result.Error
	}
	plan.LectureID = lectureID
	plan.UserID = userID
	plan.StartTime = startTime
	plan.EndTime = endTime
	plan.Status = status
	result = pr.DB.Save(plan)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *PlanRepository) DeletePlanByID(id int) error {
	result := pr.DB.Unscoped().Delete(&models.Plan{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
