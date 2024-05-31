package requests

import "time"

type CreatePlanRequest struct {
	LectureName string    `json:"lecture_name"`
	UserID      uint      `json:"user_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
