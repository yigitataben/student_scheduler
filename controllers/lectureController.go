package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/yigitataben/student_scheduler/models"
	"github.com/yigitataben/student_scheduler/services"
	"net/http"
)

type LectureController struct {
	LectureService services.LectureService
}

func NewLectureController(lectureService *services.LectureService) *LectureController {
	return &LectureController{LectureService: *lectureService}
}

func (lc *LectureController) CreateLectures(c echo.Context) error {
	type LectureInput struct {
		LectureName string `json:"LectureName"`
	}

	var lecturesInput []LectureInput
	if err := c.Bind(&lecturesInput); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to read body"})
	}

	var lectures []models.Lecture
	for i, input := range lecturesInput {
		lecture := models.Lecture{
			LectureName: input.LectureName,
			LectureID:   uint(i + 1),
		}
		lectures = append(lectures, lecture)
	}

	err := lc.LectureService.CreateLectures(lectures)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create lectures"})
	}

	return c.JSON(http.StatusOK, echo.Map{})
}

func (lc *LectureController) GetAllLectures(c echo.Context) error {
	lectures, err := lc.LectureService.GetAllLectures()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve lectures"})
	}
	return c.JSON(http.StatusOK, echo.Map{"lectures": lectures})
}

func (lc *LectureController) GetLecture(c echo.Context) error {
	ID := c.Param("id")
	lecture, err := lc.LectureService.GetLecture(ID)
	if err != nil {
		if errors.Is(err, services.ErrLectureNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Lecture not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve lecture"})
	}
	return c.JSON(http.StatusOK, echo.Map{"lecture": lecture})
}

func (lc *LectureController) DeleteLecture(c echo.Context) error {
	ID := c.Param("id")
	err := lc.LectureService.DeleteLecture(ID)
	if err != nil {
		if errors.Is(err, services.ErrLectureNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Lecture not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete lecture"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Lecture deleted successfully"})
}
