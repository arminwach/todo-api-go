package services

import (
	"errors"
	"todo-api/db"
	"todo-api/models"

	"database/sql"

	_ "modernc.org/sqlite"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
	ErrTaskEmpty    = errors.New("task cannot be empty")
)

func CreateTodo(task string) (models.Todo, error) {
	if task == "" {
		return models.Todo{}, ErrTaskEmpty
	}

	result, err := db.DB.Exec("INSERT INTO todos(task) VALUES(?)", task)
	if err != nil {
		return models.Todo{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Todo{}, err
	}

	return models.Todo{
		ID:   int(id),
		Task: task,
	}, nil
}

func GetTodoByID(id int) (models.Todo, error) {
	var todo models.Todo
	if err := db.DB.QueryRow("SELECT id, task FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Task); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Todo{}, ErrTodoNotFound
		}
		return models.Todo{}, err

	}

	return todo, nil
}

func GetAllTodos() ([]models.Todo, error) {
	rows, err := db.DB.Query("SELECT id, task FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []models.Todo{}

	for rows.Next() {
		todo := models.Todo{}

		if err := rows.Scan(&todo.ID, &todo.Task); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func UpdateTodo(id int, task string) (models.Todo, error) {
	if task == "" {
		return models.Todo{}, ErrTaskEmpty
	}

	result, err := db.DB.Exec("UPDATE todos SET task = ? WHERE id = ?", task, id)
	if err != nil {
		return models.Todo{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Todo{}, err
	}

	if rowsAffected == 0 {
		return models.Todo{}, ErrTodoNotFound
	}

	return GetTodoByID(id)
}

func DeleteTodo(id int) error {
	result, err := db.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrTodoNotFound
	}

	return nil
}
