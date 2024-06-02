package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yigitataben/student_scheduler/requests"
	"github.com/yigitataben/student_scheduler/services"
)

type LectureController struct {
	LectureService *services.LectureService
}

func NewLectureController(lectureService *services.LectureService) *LectureController {
	return &LectureController{LectureService: lectureService}
}

func (lc *LectureController) CreateLectures(c echo.Context) error {
	var lectureRequests []requests.CreateLectureRequest
	if err := c.Bind(&lectureRequests); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to read body"})
	}
	err := lc.LectureService.CreateLectures(lectureRequests)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to create lectures"})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "Lectures created successfully"})
}

func (lc *LectureController) GetAllLectures(c echo.Context) error {
	lectures, err := lc.LectureService.GetAllLectures()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch lectures"})
	}
	return c.JSON(http.StatusOK, lectures)
}

func (lc *LectureController) GetLectureByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := lc.LectureService.GetLectureByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Lecture not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (lc *LectureController) UpdateLectureByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	var updateLecture struct {
		LectureName string `json:"lecture_name"`
	}
	if err := c.Bind(&updateLecture); err != nil {
		return err
	}

	if err := lc.LectureService.UpdateLectureByID(id, updateLecture.LectureName); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(http.StatusOK, "Lecture updated successfully")
}

func (lc *LectureController) DeleteLectureByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	if err := lc.LectureService.DeleteLectureByID(id); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.JSON(http.StatusOK, "Lecture deleted successfully")
}
