package main

import (
	"net/http"
	"todo-api/routes"
)

func main() {
	r := routes.SetupRouter()
	http.ListenAndServe(":8080", r)
}
