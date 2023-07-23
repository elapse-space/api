package models

import "api/pkg/common/db/models"

type User struct {
	Id    int32  `json:"id" gorm:"primary_key;auto_increment"`
	Name  string `json:"name" gorm:"unique"`
	Price string `json:"price"`
	models.Model
}
