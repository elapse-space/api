package db

import (
	"api/pkg/common/config"
	"api/pkg/common/utils"
	"api/pkg/domain/user/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(c *config.Config) (db *gorm.DB, err error) {
	defer func() {
		if err == nil {
			utils.Logger.Info("DB initiated successfully")
		}
	}()

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return
	}
	err = db.AutoMigrate(&models.User{})
	return
}
