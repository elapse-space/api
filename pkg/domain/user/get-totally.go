package user

import (
	"api/pkg/common/utils"
	"api/pkg/domain/user/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUsersTotally(c *fiber.Ctx) error {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.BaseResponse{
			Data:    nil,
			Message: "Error getting information about totally users",
		})
	}

	response := make([]fiber.Map, 0, len(users))
	for _, user := range users {
		userData := fiber.Map{
			"id":         user.Id,
			"email":      user.Email,
			"username":   user.Username,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
			"deleted_at": user.DeletedAt,
		}
		response = append(response, userData)
	}

	return c.Status(fiber.StatusOK).JSON(utils.BaseResponse{
		Data:    response,
		Message: "Information about total users",
	})
}
