package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Middleware(c *fiber.Ctx) error {
	jwt := c.Cookies("token")
	if jwt == "" {
		return fiber.NewError(409, "Отказано в доступе")
	}
	return c.Next()
}
