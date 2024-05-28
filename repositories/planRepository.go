package repositories

import (
	"errors"
	"github.com/yigitataben/student_scheduler/models"
	"gorm.io/gorm"
)

type PlanRepository struct {
	DB *gorm.DB
}

func NewPlanRepository(db *gorm.DB) *PlanRepository {
	return &PlanRepository{DB: db}
}

func (r *PlanRepository) FindConflictingPlans(userID uint, startTime, endTime int64) ([]models.Plan, error) {
	var existingPlans []models.Plan
	err := r.DB.Where("student_id = ? AND start_time < ? AND end_time > ?", userID, endTime, startTime).Find(&existingPlans).Error
	return existingPlans, err
}

func (r *PlanRepository) Create(plan *models.Plan) error {
	return r.DB.Create(plan).Error
}

func (r *PlanRepository) FindByID(id string) (*models.Plan, error) {
	var plan models.Plan
	err := r.DB.First(&plan, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &plan, err
}

func (r *PlanRepository) Save(plan *models.Plan) error {
	return r.DB.Save(plan).Error
}

func (r *PlanRepository) FindAll() ([]models.Plan, error) {
	var plans []models.Plan
	err := r.DB.Order("created_at desc").Find(&plans).Error
	return plans, err
}

func (r *PlanRepository) Delete(plan *models.Plan) error {
	return r.DB.Unscoped().Delete(plan).Error
}
