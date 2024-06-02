package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yigitataben/student_scheduler/services"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uc *UserController) SignUp(c echo.Context) error {
	var userSignUpRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&userSignUpRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to read body"})
	}
	err := uc.UserService.SignUp(userSignUpRequest.Email, userSignUpRequest.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to sign up user"})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "User signed up successfully"})
}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch users"})
	}
	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := uc.UserService.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	var updateUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&updateUser); err != nil {
		return err
	}

	if err := uc.UserService.UpdateUserByID(id, updateUser.Email, updateUser.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(http.StatusOK, "User updated successfully")
}

func (uc *UserController) DeleteUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	if err := uc.UserService.DeleteUserByID(id); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.JSON(http.StatusOK, "User deleted successfully")
}
