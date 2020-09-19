package todo

import (
	"strconv"

	"github.com/gofiber/fiber"
)

var todos []Todo

func init() {
	todos = append(todos, Todo{Title: "Primeiro", Checked: true})
	todos = append(todos, Todo{Title: "Segundo", Checked: false})
}

//All list all todos
func All(c *fiber.Ctx) {
	c.JSON(todos)
	c.Append("content-type", "application/json")
	c.SendStatus(200)
}

//Get show a todo
func Get(c *fiber.Ctx) {
	c.Append("content-type", "application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id >= len(todos) {
		c.JSON(&fiber.Map{
			"ok":    false,
			"error": "not found",
		})
		c.SendStatus(404)
		return
	}

	todo := todos[id]
	c.JSON(todo)
	c.SendStatus(200)
}

//Add create a todo
func Add(c *fiber.Ctx) {
	c.Append("content-type", "application/json")
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		c.Status(400).Send(err)
		return
	}
	todos = append(todos, *todo)
	c.JSON(todo)
	c.SendStatus(200)
}

//Update modify a todo
func Update(c *fiber.Ctx) {
	c.Append("content-type", "application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id >= len(todos) {
		c.Status(404).Send("not found")
		return
	}
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		c.Status(400).Send("bad request")
		return
	}
	todos[id] = *todo
	c.JSON(todo)
	c.SendStatus(200)
}

//Destroy remove a todo
func Destroy(c *fiber.Ctx) {
	c.Append("content-type", "application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id >= len(todos) {
		c.Status(404).Send("not found")
		return
	}
	todos = append(todos[:id], todos[id+1:]...)
	c.JSON(&fiber.Map{
		"ok": true,
	})
	c.SendStatus(200)
}
