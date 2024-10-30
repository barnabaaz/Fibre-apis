package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteTodo(c *fiber.Ctx) error {
	params := c.Params("id")

	i, err := strconv.Atoi(params)
	if err != nil {
		return c.SendString("no valid item to delete")
	}
	DeleteTodoMethod(i)
	return c.JSON(Todos)
}
