package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos = []Todo{
	{ID: 1, Task: "Clean Code"},
	{ID: 2, Task: "Eat"},
}

func getIDFromPath(path string) (int, error) {
	parts := strings.Split(path, "/")
	return strconv.Atoi(parts[len(parts)-1])
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo

	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updated Todo
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Task = updated.Task

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func main() {
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
