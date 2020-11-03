package models

type User struct {
	ID       uint64
	Login string `json:"-" validate:"required,email"`
	Password string `json:"-"`
}

type AuthInput struct {
	Login    string `validate:"required,email"`
	Password string
}

type RegistrationInput struct {
	Login    string `validate:"required,email"`
	Password string
	Name	 string
	Surname	 string
}
