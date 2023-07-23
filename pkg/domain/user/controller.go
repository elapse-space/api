package user

import (
	"api/pkg/common/utils"
	"api/pkg/domain/auth"
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
	routes.Get("/totally", auth.MiddleWare(), h.GetUsersTotally)
	routes.Get("/me", auth.MiddleWare(), h.GetUserMe)

	utils.Logger.Info("Routes registered successfully")
}
