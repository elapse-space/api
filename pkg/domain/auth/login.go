package auth

import (
	"api/pkg/common/utils"
	model "api/pkg/domain/auth/models"
	"api/pkg/domain/user/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) Login(c *fiber.Ctx) error {
	var body model.Login
	validate := validator.New()

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BaseResponse{
			Data:    nil,
			Message: "Invalid body payload",
		})
	}

	if err := validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BaseResponse{
			Data:    nil,
			Message: "Validation error: " + err.Error(),
		})
	}

	var user models.User
	if result := h.DB.Where("email = ?", body.Email).First(&user); result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.BaseResponse{
			Data:    nil,
			Message: "Invalid email or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.BaseResponse{
			Data:    nil,
			Message: "Invalid email or password",
		})
	}

	accessToken, err := utils.GenerateJWTToken(user.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BaseResponse{
			Data:    nil,
			Message: "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utils.BaseResponse{
		Data: fiber.Map{
			"user": fiber.Map{
				"id":       &user.Id,
				"email":    &user.Email,
				"username": &user.Username,
			},
			"access_token": accessToken,
		},
		Message: "User successfully logged in",
	})
}
