package register_handler

type RegisterDtoIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
