package models

import "api/pkg/common/db/models"

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	models.Model
}
