package todo

import (
	"errors"
	"todo_list/repository"

	"gorm.io/gorm"
)

//ListTodos list all todos
func ListTodos() ([]Todo, error) {
	var todos []Todo
	db := repository.GetDB()
	result := db.Find(&todos)
	return todos, result.Error
}

//CreateTodo take a todo
func CreateTodo(todo *Todo) error {
	db := repository.GetDB()
	result := db.Create(todo)
	return result.Error
}

//GetTodo take a todo
func GetTodo(id int) (Todo, error) {
	var todo Todo
	db := repository.GetDB()
	result := db.First(&todo, id)
	return todo, result.Error
}

//UpdateTodo change a todo
func UpdateTodo(id int, params *Todo) (Todo, error) {
	var todo Todo
	db := repository.GetDB()
	if result := db.First(&todo, id); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return todo, gorm.ErrRecordNotFound
	}
	result := db.Model(&todo).Omit("id").Updates(params)
	return todo, result.Error
}

//DestroyTodo remove a todo
func DestroyTodo(id int) error {
	var todo Todo
	db := repository.GetDB()
	if result := db.First(&todo, id); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	db.Delete(&todo, id)
	return nil
}
