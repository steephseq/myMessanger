package models

type UserLogin struct {
	Email    string
	Password string
}

type UserRegister struct {
	Email    string
	Password string
	Name     string
	Username string
}
