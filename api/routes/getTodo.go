package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetTodo(c *fiber.Ctx) error {
	fmt.Println(Todos)
	return c.JSON(Todos)
}
