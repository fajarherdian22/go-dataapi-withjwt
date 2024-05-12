package domain

type User struct {
	Id       int
	Username string
	Password string
	Email    string
	Auths    string
}

type CreateUser struct {
	Username string
	Password string
	Email    string
}
