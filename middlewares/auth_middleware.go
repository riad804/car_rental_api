package middlewares

import (
	"os"

	"github.com/car_rental_api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET") // In production, use environment variable
}

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claims, err := utils.ParseJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	c.Locals("userID", uint(claims.UserID))

	return c.Next()
}
