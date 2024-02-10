package register_service

type RegisterDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
