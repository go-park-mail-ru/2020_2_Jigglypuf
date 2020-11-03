package models

type Profile struct {
	UserModel  *User `json:"-"`
	Name       string
	Surname    string
	AvatarPath string
}

type ProfileFormData struct {
	Name    string
	Surname string
	Avatar  string
}
