package models

import "github.com/gofiber/fiber/v2"

type Response[T any] struct {
	Status  bool   `json:"status"`         // "success" or "error"
	Message string `json:"message"`        // Summary or description
	Data    *T     `json:"data,omitempty"` // Payload (can be null)
}

func Success[T any](c *fiber.Ctx, message string, data T) error {
	res := Response[T]{
		Status:  true,
		Message: message,
		Data:    &data,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func Error(c *fiber.Ctx, code int, message string) error {
	res := Response[any]{
		Status:  false,
		Message: message,
	}
	return c.Status(code).JSON(res)
}
