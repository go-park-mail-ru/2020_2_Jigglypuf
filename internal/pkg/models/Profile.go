package models

type Profile struct {
	Login      *User `json:"-"`
	Name       string
	Surname    string
	AvatarPath string
}
