package auth

import (
	"api/pkg/common/utils"
	model "api/pkg/domain/auth/models"
	"api/pkg/domain/user/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) Registration(c *fiber.Ctx) error {
	var body model.Registration
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
	if result := h.DB.Where("email = ?", body.Email).First(&user); result.Error == nil {
		return c.Status(fiber.StatusConflict).JSON(utils.BaseResponse{
			Data:    nil,
			Message: "Email already in use",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BaseResponse{
			Data:    nil,
			Message: "Failed to hash password",
		})
	}

	newUser := models.User{
		Email:    body.Email,
		Username: body.Username,
		Password: string(hashedPassword),
	}

	if result := h.DB.Create(&newUser); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BaseResponse{
			Data:    nil,
			Message: "Failed to create user",
		})
	}

	accessToken, _ := utils.GenerateJWTToken(newUser.Id)

	return c.Status(fiber.StatusCreated).JSON(utils.BaseResponse{
		Data: fiber.Map{
			"user": fiber.Map{
				"id":       &newUser.Id,
				"email":    &newUser.Email,
				"username": &newUser.Username,
			},
			"access_token": accessToken,
		},
		Message: "User successfully registered",
	})
}
