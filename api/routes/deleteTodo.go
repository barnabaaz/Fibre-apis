package routes

import "github.com/gofiber/fiber/v2"

func DeleteTodo(c *fiber.Ctx) error {
	return c.SendString("Welcome from Delete")
}
