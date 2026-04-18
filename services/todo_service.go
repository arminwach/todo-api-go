package services

import (
	"errors"
	"todo-api/models"
)

func CreateTodo(task string) (models.Todo, error) {
	if task == "" {
		return models.Todo{}, errors.New("task cannot be empty")
	}

	todo := models.Todo{
		ID:   len(models.Todos) + 1,
		Task: task,
	}

	models.Todos = append(models.Todos, todo)

	return todo, nil
}

func GetTodoByID(id int) (models.Todo, error) {
	for _, todo := range models.Todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return models.Todo{}, errors.New("todo not found")
}

func UpdateTodo(id int, task string) (models.Todo, error) {
	if task == "" {
		return models.Todo{}, errors.New("task cannot be empty")
	}
	for i, todo := range models.Todos {
		if todo.ID == id {
			models.Todos[i].Task = task
			return models.Todos[i], nil
		}
	}
	return models.Todo{}, errors.New("todo not found")
}

func DeleteTodo(id int) error {
	for i, todo := range models.Todos {
		if todo.ID == id {
			models.Todos = append(models.Todos[:i], models.Todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}

func GetAllTodos() []models.Todo {
	return models.Todos
}
