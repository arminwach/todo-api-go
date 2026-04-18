package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"todo-api/services"

	"github.com/go-chi/chi/v5"
)

type CreateTodoRequest struct {
	Task string `json:"task"`
}

type UpdateTodoRequest struct {
	Task string `json:"task"`
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func handleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, services.ErrTodoNotFound):
		writeJSONError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, services.ErrTaskEmpty):
		writeJSONError(w, http.StatusBadRequest, err.Error())
	default:
		writeJSONError(w, http.StatusInternalServerError, "internal server error")
	}
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := services.GetAllTodos()
	if err != nil {
		handleServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req CreateTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	todo, err := services.CreateTodo(req.Task)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, todo)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	todo, err := services.GetTodoByID(id)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req UpdateTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	todo, err := services.UpdateTodo(id, req.Task)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := services.DeleteTodo(id); err != nil {
		handleServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
