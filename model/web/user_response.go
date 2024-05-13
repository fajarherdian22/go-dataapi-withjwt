package web

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Auths    string `json:"auths"`
}
