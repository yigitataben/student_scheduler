package requests

type CreateLectureRequest struct {
	ID          int    `json:"id"`
	LectureName string `json:"lecture_name"`
}
