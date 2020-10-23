package models

type User struct {
	ID       uint64
	Username string `json:"-"`
	Password string `json:"-"`
}

type AuthInput struct {
	Login    string
	Password string
}

type RegistrationInput struct {
	Login    string
	Password string
}
