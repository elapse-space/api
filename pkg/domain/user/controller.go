package user

import (
	"api/pkg/common/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func Controller(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/user")
	routes.Get("/totally", h.GetUsersAll)

	utils.Logger.Info("Routes registered successfully")
}
