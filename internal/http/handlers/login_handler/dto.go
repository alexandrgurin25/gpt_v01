package login_handler

type LoginDtoIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDtoOut struct {
	AccessToken string `json:"accessToken"`
}
