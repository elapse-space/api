package user

import (
	"api/pkg/common/utils"
	"api/pkg/domain/user/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUsersAll(c *fiber.Ctx) error {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		utils.Logger.Error(result.Error.Error())
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	responseData := make([]fiber.Map, 0, len(users))
	for _, user := range users {
		userData := fiber.Map{
			"id":         user.Id,
			"username":   user.Username,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
			"deleted_at": user.DeletedAt,
		}
		responseData = append(responseData, userData)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(utils.BaseResponse{Data: responseData, Message: "Information about total users", Success: true})
}
