package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yigitataben/student_scheduler/controllers"
)

func LectureRoutes(e *echo.Echo, lectureController *controllers.LectureController) {
	e.POST("/lectures", lectureController.CreateLectures)
	e.GET("/lectures", lectureController.GetAllLectures)
}
