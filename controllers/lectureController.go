package controllers

import (
	"errors"
	"net/http"

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

func (lc *LectureController) GetLecture(c echo.Context) error {
	id := c.Param("id")
	lecture, err := lc.LectureService.GetLecture(id)
	if err != nil {
		if errors.Is(err, services.ErrLectureNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Lecture not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch lecture"})
	}
	return c.JSON(http.StatusOK, lecture)
}

func (lc *LectureController) DeleteLecture(c echo.Context) error {
	id := c.Param("id")
	err := lc.LectureService.DeleteLecture(id)
	if err != nil {
		if errors.Is(err, services.ErrLectureNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Lecture not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete lecture"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Lecture deleted successfully"})
}
