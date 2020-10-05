package models

type Profile struct {
	Login *User
	Name string
	Surname string
	AvatarPath string
}
