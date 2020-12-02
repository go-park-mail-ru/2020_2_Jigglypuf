package models

type Profile struct {
	UserCredentials *User
	Name            string
	Surname         string
	AvatarPath      string
}

type ProfileFormData struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Avatar  string `json:"avatar"`
}
