package models

type Profile struct {
	UserCredentials *User
	Name            string
	Surname         string
	AvatarPath      string
}

type ProfileFormData struct {
	Name    string
	Surname string
	Avatar  string
}
