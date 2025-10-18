package middlewares

import (
	"strings"

	"github.com/Auxesia23/velarsyapi/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func JWTAuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.ErrUnauthorized
	}
	if ok := strings.HasPrefix(authHeader, "Bearer"); !ok {
		return fiber.ErrUnauthorized
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return fiber.ErrUnauthorized
	}
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return fiber.ErrUnauthorized
	}
	c.Locals("user", claims)
	return c.Next()
}
