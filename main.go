package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yigitataben/student_scheduler/controllers"
	"github.com/yigitataben/student_scheduler/initializers"
	"github.com/yigitataben/student_scheduler/repositories"
	"github.com/yigitataben/student_scheduler/routes"
	"github.com/yigitataben/student_scheduler/services"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()

	e := echo.New()

	userRepository := repositories.NewUserRepository(initializers.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	lectureRepository := repositories.NewLectureRepository(initializers.DB)
	lectureService := services.NewLectureService(lectureRepository)
	lectureController := controllers.NewLectureController(lectureService)

	planRepository := repositories.NewPlanRepository(initializers.DB)
	planService := services.NewPlanService(planRepository)
	planController := controllers.NewPlanController(planService)

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	}))

	// Register routes
	routes.UserRoutes(e, userController)
	routes.LectureRoutes(e, lectureController)
	routes.PlanRoutes(e, planController)

	// Start server
	port := os.Getenv("PORT")
	if err := e.Start(":" + port); err != nil {
		panic("Failed to start server")
	}
}
