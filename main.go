package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	Id        int `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}



func main() {

	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {	
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Create a Todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {	
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == ""{
			return c.Status(400).JSON(fiber.Map{"Error":"No body"})
		}
		todo.Id = len(todos) + 1
		todos = append(todos, *todo)
		return c.Status(200).JSON(todo)
	})


	// Update a todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error{
		id := c.Params("id")
		for i, todo := range todos{
			if fmt.Sprint(todo.Id) == id{
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"Error": "Todo not found"})
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx)error{
		id := c.Params("id")
		for i, todo := range todos{
			if fmt.Sprint(todo.Id) == id{
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"Success": "True"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"Error": "Todo not found"})
	})

	log.Fatal(app.Listen(":"+ PORT))
}
