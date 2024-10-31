package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/barnabaaz/Fibre-apis/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

//Database Instance

var db *sql.DB

//Database settings

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "barnabas"
)

func connnectDb() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil

}

func setupRoutes(app *fiber.App) {
	app.Get("/api/todo", routes.GetAllTdo)
	app.Post("/api/todo", routes.CreateTodo)
	app.Delete("/api/todo/:id", routes.DeleteTodo)
	//get single Todo Item
	app.Get("/api/todo/:id", routes.GetTodo)
	//Update Todo
	app.Patch("/api/todo/:id", routes.EditTodo)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Word")
	})
}
func main() {
	if err := connnectDb(); err != nil {
		log.Fatal(err)

	}
	app := fiber.New()
	//get all todo items
	app.Use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
	fmt.Println("Server is starting on port 3000")

}
