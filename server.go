package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos = []Todo{}

func main() {

	app := fiber.New()

	app.Get("/", getAll)

	app.Post("/", create)

	app.Put("/:id", update)

	app.Delete("/:id", delete)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

}

func getAll(c *fiber.Ctx) error {
	return c.Status(200).JSON(todos)
}

func create(c *fiber.Ctx) error {

	newTodo := Todo{}

	if err := c.BodyParser(&newTodo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "json is not valid",
		})
	}

	todos = append(todos, newTodo)

	return c.SendStatus(201)
}

func update(c *fiber.Ctx) error {
	return c.SendString("WIP")
}

func delete(c *fiber.Ctx) error {

	id := c.Params("id")
	newTodos := []Todo{}

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "type the id",
		})
	}

	i, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "type valid id",
		})
	}

	for _, v := range todos {
		if v.Id != int(i) {
			newTodos = append(newTodos, v)
		}
	}

	todos = newTodos

	return c.SendStatus(204)
}
