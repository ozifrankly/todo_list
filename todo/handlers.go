package todo

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//All list all todos
func All(c *fiber.Ctx) error {
	todos, err := ListTodos()
	if err != nil {
		return errorMessage(c, 500, err.Error())
	}
	c.JSON(todos)
	return c.SendStatus(200)
}

//Get show a todo
func Get(c *fiber.Ctx) error {
	c.Append("content-type", "application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errorMessage(c, 400, "bad request")
	}

	todo, err := GetTodo(id)
	if err != nil {
		return errorMessage(c, 404, "not found")
	}
	c.JSON(todo)
	return c.SendStatus(200)
}

//Add create a todo
func Add(c *fiber.Ctx) error {
	c.Append("content-type", "application/json")
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return errorMessage(c, 400, "bad request")
	}

	if err := CreateTodo(todo); err != nil {
		return errorMessage(c, 304, err.Error())
	}
	c.JSON(todo)
	return c.SendStatus(201)
}

//Update modify a todo
func Update(c *fiber.Ctx) error {
	c.Append("content-type", "application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errorMessage(c, 404, "not found")
	}
	params := new(Todo)
	if err := c.BodyParser(params); err != nil {
		return errorMessage(c, 400, "bad request")
	}
	fmt.Println(params)
	todo, err := UpdateTodo(id, params)
	if err != nil {
		return errorMessage(c, 400, "bad request")
	}
	c.JSON(todo)
	return c.SendStatus(200)
}

//Destroy remove a todo
func Destroy(c *fiber.Ctx) error {
	c.Append("content-type", "application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errorMessage(c, 404, "not found")
	}
	err = DestroyTodo(id)
	if err != nil {
		return errorMessage(c, 404, "not found")
	}
	c.JSON(&fiber.Map{
		"ok": true,
	})
	return c.SendStatus(200)
}

func errorMessage(c *fiber.Ctx, code int, message string) error {
	c.JSON(&fiber.Map{
		"ok":    false,
		"error": message,
	})
	return c.SendStatus(code)
}
