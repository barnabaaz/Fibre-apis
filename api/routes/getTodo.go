package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAllTdo(c *fiber.Ctx) error {
	fmt.Println(Todos)
	return c.JSON(Todos)
}

func GetTodo(c *fiber.Ctx) error {
	params := c.Params("id")
	i, err := strconv.Atoi(params)
	if err != nil {
		return c.SendString("this is not allowed")
	}
	for _, item := range Todos {
		if item.ID == i {
			return c.JSON(item)
		}
	}
	return c.JSON("")
}
