package models

import "2020_2_Jigglypuf/backup/models"

type Profile struct {
	Login *models.User
	Name string
	Surname string
	AvatarPath string
}
