package main

import (
	"log"
	"todo_list/middleware"
	"todo_list/repository"
	"todo_list/todo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := repository.Connect()
	db.Debug().AutoMigrate(&todo.Todo{})
}

func main() {
	app := fiber.New()
	app.Use(middleware.Header)
	app.Use(cors.New())

	v1 := app.Group("api/v1")
	v1.Get("/todos", todo.All)
	v1.Post("/todos", todo.Add)
	v1.Get("/todos/:id", todo.Get)
	v1.Put("/todos/:id", todo.Update)
	v1.Delete("/todos/:id", todo.Destroy)
	app.Listen(":3000")
}
