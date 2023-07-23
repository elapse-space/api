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
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(utils.BaseResponse{Data: nil, Message: "Unauthorized", Success: false})
		}

		tokenString := authHeader[7:]
		if tokenString == "" {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(utils.BaseResponse{Data: nil, Message: "Unauthorized", Success: false})
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Status(fiber.StatusUnauthorized)
				return c.JSON(utils.BaseResponse{Data: nil, Message: "Unauthorized", Success: false})
			}
			c.Status(fiber.StatusBadRequest)
			return c.JSON(utils.BaseResponse{Data: nil, Message: "Bad Request", Success: false})
		}

		if !token.Valid {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(utils.BaseResponse{Data: nil, Message: "Unauthorized", Success: false})
		}

		claims, ok := token.Claims.(*model.Claims)
		if !ok {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(utils.BaseResponse{Data: nil, Message: "Internal Server Error", Success: false})
		}
		c.Locals("user", claims)

		return c.Next()
	}
}
