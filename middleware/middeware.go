package middleware

import (
	"github.com/gofiber/fiber/v2"
)

//Header set http header
func Header(c *fiber.Ctx) error {
	c.Append("content-type", "application/json")
	return c.Next()
}
