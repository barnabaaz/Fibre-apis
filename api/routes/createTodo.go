package routes

import (
	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	todo := new(Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return err
	}
	todos := AddTododMethod(Todo{ID: todo.ID, DESCRIPTION: todo.DESCRIPTION, STATUS: todo.STATUS})

	return c.JSON(todos)
}
