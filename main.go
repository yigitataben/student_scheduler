package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/yigitataben/student_scheduler/initializers"
	"github.com/yigitataben/student_scheduler/routes"
	"os"
)

func init() {
	err := initializers.LoadEnvVariables()
	if err != nil {
		return
	}
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	e := echo.New()

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	}))

	// Register routes
	routes.UserRoutes(e)
	routes.LectureRoutes(e)
	routes.PlanRoutes(e)

	// Start server
	port := os.Getenv("PORT")
	if err := e.Start(":" + port); err != nil {
		panic("Failed to start server")
	}
}
