package requests

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
