package todo

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
)

//All list all todos
func All(c *fiber.Ctx) {
	c.Append("content-type", "application/json")
	todos, err := ListTodos()
	if err != nil {
		errorMessage(c, 500, err.Error())
		return
	}
	c.JSON(todos)
	c.SendStatus(200)
}

//Get show a todo
func Get(c *fiber.Ctx) {
	c.Append("content-type", "application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		errorMessage(c, 400, "bad request")
		return
	}

	todo, err := GetTodo(id)
	if err != nil {
		errorMessage(c, 404, "not found")
		return
	}
	c.JSON(todo)
	c.SendStatus(200)
}

//Add create a todo
func Add(c *fiber.Ctx) {
	c.Append("content-type", "application/json")
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		errorMessage(c, 400, "bad request")
		return
	}

	if err := CreateTodo(todo); err != nil {
		errorMessage(c, 304, err.Error())
		return
	}
	c.JSON(todo)
	c.SendStatus(201)
}

//Update modify a todo
func Update(c *fiber.Ctx) {
	c.Append("content-type", "application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		errorMessage(c, 404, "not found")
		return
	}
	params := new(Todo)
	if err := c.BodyParser(params); err != nil {
		errorMessage(c, 400, "bad request")
		return
	}
	fmt.Println(params)
	todo, err := UpdateTodo(id, params)
	if err != nil {
		errorMessage(c, 400, "bad request")
		return
	}
	c.JSON(todo)
	c.SendStatus(200)
}

//Destroy remove a todo
func Destroy(c *fiber.Ctx) {
	c.Append("content-type", "application/json")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		errorMessage(c, 404, "not found")
		return
	}
	err = DestroyTodo(id)
	if err != nil {
		errorMessage(c, 404, "not found")
		return
	}
	c.JSON(&fiber.Map{
		"ok": true,
	})
	c.SendStatus(200)
}

func errorMessage(c *fiber.Ctx, code int, message string) {
	c.JSON(&fiber.Map{
		"ok":    false,
		"error": message,
	})
	c.SendStatus(code)
}
