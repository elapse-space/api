package auth

import (
	"api/pkg/common/utils"
	model "api/pkg/domain/auth/models"
	"api/pkg/domain/user/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) Registration(c *fiber.Ctx) error {
	var body model.Registration

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(utils.BaseResponse{Data: nil, Message: "Invalid body payload", Success: false})
	}

	var user models.User
	if result := h.DB.Where("username = ?", body.Username).First(&user); result.Error != nil {
	} else {
		c.Status(fiber.StatusConflict)
		return c.JSON(utils.BaseResponse{Data: nil, Message: "Username already in use", Success: false})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(utils.BaseResponse{Data: nil, Message: "Failed to hash password", Success: false})
	}

	newUser := models.User{
		Username: body.Username,
		Password: string(hashedPassword),
	}

	if result := h.DB.Create(&newUser); result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(utils.BaseResponse{Data: nil, Message: "Failed to create user", Success: false})
	}

	token, _ := utils.GenerateJWTToken(newUser.Username)

	c.Status(fiber.StatusCreated)
	return c.JSON(utils.BaseResponse{Data: fiber.Map{"access_token": token}, Message: "User successfully registered", Success: true})
}
