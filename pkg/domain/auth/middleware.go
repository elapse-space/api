package auth

import (
	"api/pkg/common/utils"
	model "api/pkg/domain/auth/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtSecret = []byte("g78asdg78ags87g89hhasdj")

func MiddleWare() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.BaseResponse{
				Data:    nil,
				Message: "Unauthorized",
			})
		}

		tokenString := authHeader[7:]
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.BaseResponse{
				Data:    nil,
				Message: "Unauthorized",
			})
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.Status(fiber.StatusUnauthorized).JSON(utils.BaseResponse{
					Data:    nil,
					Message: "Unauthorized",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(utils.BaseResponse{
				Data:    nil,
				Message: "Bad Request",
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.BaseResponse{
				Data:    nil,
				Message: "Unauthorized",
			})
		}

		claims, ok := token.Claims.(*model.Claims)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.BaseResponse{
				Data:    nil,
				Message: "Internal Server Error",
			})
		}
		c.Locals("user", claims)

		return c.Next()
	}
}
