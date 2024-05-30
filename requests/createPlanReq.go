package requests

type CreatePlanRequest struct {
	LectureName string `json:"lecture_name"`
	UserID      uint   `json:"user_id"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}
