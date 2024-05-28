package controllers

import (
	"github.com/yigitataben/student_scheduler/models"
	"github.com/yigitataben/student_scheduler/services"
	"net/http"
)

type PlanHandler struct {
	PlanService services.PlanService
}

func NewPlanHandler(planService services.PlanService) *PlanHandler {
	return &PlanHandler{PlanService: planService}
}

func (h *PlanHandler) CreatePlan(c echo.Context) error {
	plan := new(models.Plan)
	if err := c.Bind(plan); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.PlanService.CreatePlan(plan)
	if err != nil {
		if err == services.ErrSchedulingConflict {
			return c.JSON(http.StatusConflict, "Scheduling conflict")
		}
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusCreated, plan)
}

func (h *PlanHandler) UpdatePlanStatus(c echo.Context) error {
	id := c.Param("id")
	status := c.QueryParam("status")

	plan, err := h.PlanService.UpdatePlanStatus(id, status)
	if err != nil {
		if err == services.ErrPlanNotFound {
			return c.JSON(http.StatusNotFound, "Plan not found")
		}
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, plan)
}

func (h *PlanHandler) GetAllPlans(c echo.Context) error {
	plans, err := h.PlanService.GetAllPlans()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve plans"})
	}
	return c.JSON(http.StatusOK, echo.Map{"plans": plans})
}

func (h *PlanHandler) GetPlanByID(c echo.Context) error {
	id := c.Param("id")

	plan, err := h.PlanService.GetPlanByID(id)
	if err != nil {
		if err == services.ErrPlanNotFound {
			return c.JSON(http.StatusNotFound, "Plan not found")
		}
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, echo.Map{"plan": plan})
}

func (h *PlanHandler) UpdatePlan(c echo.Context) error {
	id := c.Param("id")

	var newPlanData models.Plan
	if err := c.Bind(&newPlanData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to read body"})
	}

	plan, err := h.PlanService.UpdatePlan(id, &newPlanData)
	if err != nil {
		if err == services.ErrPlanNotFound {
			return c.JSON(http.StatusNotFound, "Plan not found")
		}
		return c.JSON(http.StatusInternalServerError, "Failed to update plan")
	}

	return c.JSON(http.StatusOK, echo.Map{"plan": plan})
}

func (h *PlanHandler) DeletePlan(c echo.Context) error {
	id := c.Param("id")

	err := h.PlanService.DeletePlan(id)
	if err != nil {
		if err == services.ErrPlanNotFound {
			return c.JSON(http.StatusNotFound, "Plan not found")
		}
		return c.JSON(http.StatusInternalServerError, "Failed to delete plan")
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Plan deleted successfully"})
}
