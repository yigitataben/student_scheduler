package routes

import (
	"github.com/yigitataben/student_scheduler/controllers"
)

func UserRoutes(e *echo.Echo, userController *controllers.UserController) {
	e.POST("/signup", userController.SignUp)
	e.GET("/users", userController.GetAllUsers)
	e.GET("/users/:id", userController.GetUserByID)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)
}