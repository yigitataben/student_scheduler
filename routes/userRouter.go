package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yigitataben/student_scheduler/controllers"
)

func UserRoutes(e *echo.Echo, userController *controllers.UserController) {
	e.POST("/signup", userController.SignUp)
	e.GET("/users", userController.GetAllUsers)
	e.GET("/users/:id", userController.GetUserByID)
	e.PUT("/users/:id", userController.UpdateUserByID)
	e.DELETE("/users/:id", userController.DeleteUserByID)
}
