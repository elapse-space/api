package user

import (
	"api/pkg/common/utils"
	model "api/pkg/domain/auth/models"
	"api/pkg/domain/user/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUserMe(c *fiber.Ctx) error {
	claims := c.Locals("user").(*model.Claims)
	id := claims.Id

	var user models.User
	if result := h.DB.First(&user, "id = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.BaseResponse{
			Data:    nil,
			Message: "Error getting information about own user",
		})
	}

	response := fiber.Map{
		"id":         user.Id,
		"email":      user.Email,
		"username":   user.Username,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"deleted_at": user.DeletedAt,
	}

	return c.Status(fiber.StatusOK).JSON(utils.BaseResponse{
		Data:    &response,
		Message: "Information about authorized user",
	})
}
