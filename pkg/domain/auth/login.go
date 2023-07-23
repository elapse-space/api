package auth

import (
	"api/pkg/common/utils"
	model "api/pkg/domain/auth/models"
	"api/pkg/domain/user/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) Login(c *fiber.Ctx) error {
	var body model.Login
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body payload")
	}

	var user models.User
	if result := h.DB.Where("username = ?", body.Username).First(&user); result.Error != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(utils.BaseResponse{Data: nil, Message: "Invalid username or password", Success: false})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(utils.BaseResponse{Data: nil, Message: "Invalid username or password", Success: false})
	}

	token, err := utils.GenerateJWTToken(user.Username)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(utils.BaseResponse{Data: nil, Message: "Failed to generate token", Success: false})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(utils.BaseResponse{Data: fiber.Map{"access_token": token}, Message: "User successfully login", Success: true})
}
