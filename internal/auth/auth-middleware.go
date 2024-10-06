package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Extract and validate token from request headers
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing or invalid token"})
	}

	// Validate token (JWT, etc.)
	// If valid, proceed to the next handler
	// If invalid, return unauthorized error

	return c.Next()
}
