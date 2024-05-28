package controllers

import (
	"github.com/yigitataben/student_scheduler/services"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uc *UserController) SignUp(c echo.Context) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to read body"})
	}
	err := uc.UserService.SignUp(body.Email, body.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to create user"})
	}
	return c.JSON(http.StatusOK, echo.Map{})
}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve users"})
	}
	return c.JSON(http.StatusOK, echo.Map{"users": users})
}

func (uc *UserController) GetUserByID(c echo.Context) error {
	ID := c.Param("id")
	user, err := uc.UserService.GetUserByID(ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"user": user})
}

func (uc *UserController) UpdateUser(c echo.Context) error {
	ID := c.Param("id")
	var newUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to read body"})
	}
	updatedUser, err := uc.UserService.UpdateUser(ID, newUser.Email, newUser.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update user"})
	}
	return c.JSON(http.StatusOK, echo.Map{"user": updatedUser})
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	ID := c.Param("id")
	err := uc.UserService.DeleteUser(ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "User deleted successfully"})
}
