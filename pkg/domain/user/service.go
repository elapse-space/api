package user

import (
	"api/pkg/common/utils"
	"api/pkg/domain/user/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) Welcome(c *fiber.Ctx) error {
	welcome := "Lol route"
	return c.Status(fiber.StatusOK).SendString(welcome)
}

func (h handler) GetUsers(ctx *fiber.Ctx) error {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		utils.Logger.Error(result.Error.Error())
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&users)
}
