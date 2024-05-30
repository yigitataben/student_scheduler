package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yigitataben/student_scheduler/requests"
	"github.com/yigitataben/student_scheduler/services"
)

type PlanController struct {
	PlanService *services.PlanService
}

func NewPlanController(planService *services.PlanService) *PlanController {
	return &PlanController{PlanService: planService}
}

func (pc *PlanController) CreatePlan(c echo.Context) error {
	var planRequest requests.CreatePlanRequest
	if err := c.Bind(&planRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to read body"})
	}
	err := pc.PlanService.CreatePlan(planRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to create plan"})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "Plan created successfully"})
}

func (pc *PlanController) GetAllPlans(c echo.Context) error {
	plans, err := pc.PlanService.GetAllPlans()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch plans"})
	}
	return c.JSON(http.StatusOK, plans)
}

func (pc *PlanController) GetPlanByID(c echo.Context) error {
	id := c.Param("id")
	plan, err := pc.PlanService.GetPlanByID(id)
	if err != nil {
		if errors.Is(err, services.ErrPlanNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Plan not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch plan"})
	}
	return c.JSON(http.StatusOK, plan)
}

func (pc *PlanController) UpdatePlan(c echo.Context) error {
	id := c.Param("id")
	var planRequest requests.CreatePlanRequest
	if err := c.Bind(&planRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to read body"})
	}
	plan, err := pc.PlanService.UpdatePlan(id, planRequest)
	if err != nil {
		if errors.Is(err, services.ErrPlanNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Plan not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update plan"})
	}
	return c.JSON(http.StatusOK, plan)
}

func (pc *PlanController) DeletePlan(c echo.Context) error {
	id := c.Param("id")
	err := pc.PlanService.DeletePlan(id)
	if err != nil {
		if errors.Is(err, services.ErrPlanNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Plan not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete plan"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Plan deleted successfully"})
}
