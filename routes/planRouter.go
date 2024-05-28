package routes

import (
	"github.com/yigitataben/student_scheduler/controllers"
)

func PlanRoutes(e *echo.Echo, planController *controllers.PlanController) {
	e.POST("/plans", planController.CreatePlan)
	e.GET("/plans", planController.GetAllPlans)
	e.GET("/plans/:id", planController.GetPlanByID)
	e.PUT("/plans/:id", planController.UpdatePlan)
	e.DELETE("/plans/:id", planController.DeletePlan)
}
