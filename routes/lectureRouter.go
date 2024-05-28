package routes

import (
	"github.com/yigitataben/student_scheduler/controllers"
)

func LectureRoutes(e *echo.Echo) {
	e.POST("/lectures", controllers.CreateLectures)
	e.GET("/lectures", controllers.GetAllLectures)
}
