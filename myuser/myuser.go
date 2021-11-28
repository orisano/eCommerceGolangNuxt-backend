package user

import "github.com/gofiber/fiber/v2"

func GetUser(c *fiber.Ctx) error {
	return c.SendString("hello world")
}
