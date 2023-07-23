package models

import "api/pkg/common/db/models"

type User struct {
	Id       int32  `json:"id" gorm:"primary_key;auto_increment"`
	Email    string `json:"email" gorm:"unique"`
	Username string `json:"username"`
	Password string `json:"password"`
	models.Model
}
