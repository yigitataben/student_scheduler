package requests

type CreatePlanRequest struct {
	ID        int    `json:"id"`
	LectureID int    `json:"lecture_name"`
	UserID    int    `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Status    string `json:"status"`
}
