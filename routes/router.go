package routes

import (
	"net/http"

	"todo-api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/todos", handlers.GetTodos)
	r.Post("/todos", handlers.CreateTodo)

	r.Get("/todos/{id}", handlers.GetTodoByID)
	r.Put("/todos/{id}", handlers.UpdateTodo)
	r.Delete("/todos/{id}", handlers.DeleteTodo)

	return r
}
