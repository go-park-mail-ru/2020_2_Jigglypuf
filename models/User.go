package models

import "net/http"

type User struct{
	ID uint64
	Username string
	Password string
	Cookie http.Cookie `json:"-"`
}

type AuthInput struct{
	Login string
	Password string
}

type RegistrationInput struct{
	Login string
	Password string
}

