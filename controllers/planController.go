package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yigitataben/student_scheduler/services"
)

type PlanController struct {
	PlanService *services.PlanService
}

func NewPlanController(planService *services.PlanService) *PlanController {
	return &PlanController{PlanService: planService}
}

func (pc *PlanController) CreatePlan(c echo.Context) error {
	var createPlanRequest struct {
		LectureID int    `json:"lecture_id"`
		UserID    int    `json:"user_id"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		Status    string `json:"status"`
	}
	if err := c.Bind(&createPlanRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to read body"})
	}
	err := pc.PlanService.CreatePlan(createPlanRequest.LectureID, createPlanRequest.UserID, createPlanRequest.StartTime, createPlanRequest.EndTime, createPlanRequest.Status)
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid plan ID")
	}

	plan, err := pc.PlanService.GetPlanByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Plan not found")
	}

	return c.JSON(http.StatusOK, plan)

}

func (pc *PlanController) UpdatePlan(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid plan ID")
	}

	var updatePlan struct {
		LectureID int    `json:"lecture_id"`
		UserID    int    `json:"user_id"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		Status    string `json:"status"`
	}
	if err := c.Bind(&updatePlan); err != nil {
		return err
	}

	if err := pc.PlanService.UpdatePlanByID(id, updatePlan.LectureID, updatePlan.UserID, updatePlan.StartTime, updatePlan.EndTime, updatePlan.Status); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update plan")
	}

	return c.JSON(http.StatusOK, "Plan updated successfully")

}

func (pc *PlanController) DeletePlan(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid plan ID")
	}

	if err := pc.PlanService.DeletePlanByID(id); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete plan")
	}

	return c.JSON(http.StatusOK, "Plan deleted successfully")
}
