package main

import (
	"api/pkg/common/config"
	"api/pkg/common/db"
	"api/pkg/common/utils"
	"api/pkg/domain/auth"
	"api/pkg/domain/user"
	"api/pkg/domain/user/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	utils.InitLogger()
	utils.Logger.Info("Server started")

	c, err := config.LoadConfig()
	if err != nil {
		utils.Logger.Error(err.Error())
	}

	h, err := db.Init(&c)
	if err != nil {
		utils.Logger.Error(err.Error())
	}

	err = h.AutoMigrate(&models.User{})
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	auth.Controller(app, h)
	user.Controller(app, h)

	err = app.Listen(c.Port)
	if err != nil {
		utils.Logger.Error(err.Error())
	}
}
