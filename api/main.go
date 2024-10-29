package main

import (
	"fmt"
	"log"

	"github.com/barnabaaz/Fibre-apis/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/todo", routes.GetTodo)
	app.Post("/api/todo", routes.CreateTodo)
	app.Delete("/api/todo/:id", routes.GetTodo)
	//get single Todo Item
	app.Get("/api/todo/:id", routes.GetTodo)
	//Update Todo
	app.Patch("/api/todo/:id", routes.EditTodo)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Word")
	})
}
func main() {
	app := fiber.New()
	//get all todo items
	app.Use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
	fmt.Println("Server is starting on port 3000")

}
