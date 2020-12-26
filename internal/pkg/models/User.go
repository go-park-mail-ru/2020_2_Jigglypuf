package models

type User struct {
	ID       uint64
	Login    string `validate:"required,email"`
	Password string `json:"-"`
}

type AuthInput struct {
	Login    string `validate:"required,email" json:"login"`
	Password string `json:"password"`
}

type RegistrationInput struct {
	Login    string `validate:"required,email" json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}
