package requests

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
