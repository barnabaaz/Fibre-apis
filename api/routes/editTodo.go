package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func EditTodo(c *fiber.Ctx) error {
	todo := new(Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return err
	}
	params := c.Params("id")
	i, err := strconv.Atoi(params)
	if err != nil {
		return c.SendString("Error reading values")
	}
	UpdateTodoMethod(Todo{ID: todo.ID, DESCRIPTION: todo.DESCRIPTION, STATUS: todo.STATUS}, i)
	return c.JSON(todo)
}
