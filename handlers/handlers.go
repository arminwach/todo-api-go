package handlers

import (
	"encoding/json"
	"net/http"
	"todo-api/models"
	"todo-api/utils"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo

	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newTodo.ID = len(models.Todos) + 1
	models.Todos = append(models.Todos, newTodo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, todo := range models.Todos {
		if todo.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updated models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, todo := range models.Todos {
		if todo.ID == id {
			models.Todos[i].Task = updated.Task

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Todos[i])
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, todo := range models.Todos {
		if todo.ID == id {
			models.Todos = append(models.Todos[:i], models.Todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func InitHandlers() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getTodos(w, r)
		} else if r.Method == http.MethodPost {
			createTodo(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getTodoByID(w, r)
		} else if r.Method == http.MethodPut {
			updateTodo(w, r)
		} else if r.Method == http.MethodDelete {
			deleteTodo(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
