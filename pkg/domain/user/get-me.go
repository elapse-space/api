package user

import (
	"api/pkg/common/utils"
	model "api/pkg/domain/auth/models"
	"api/pkg/domain/user/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUserMe(c *fiber.Ctx) error {
	claims := c.Locals("user").(*model.Claims)
	username := claims.Username

	var user models.User
	if result := h.DB.First(&user, "username = ?", username); result.Error != nil {
		utils.Logger.Error(result.Error.Error())
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	responseData := fiber.Map{
		"id":         user.Id,
		"username":   user.Username,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"deleted_at": user.DeletedAt,
	}

	c.Status(fiber.StatusOK)
	return c.JSON(utils.BaseResponse{Data: &responseData, Message: "Information about authorized user", Success: true})
}
