package types

type UserResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
