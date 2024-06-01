package requests

type CreatePlanRequest struct {
	LectureID uint   `json:"lecture_name"`
	UserID    uint   `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
