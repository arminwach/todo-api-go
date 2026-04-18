package main

import (
	"net/http"
	"todo-api/db"
	"todo-api/routes"
)

func main() {
	err := db.InitDB()
	if err != nil {
		panic(err)
	}
	r := routes.SetupRouter()
	http.ListenAndServe(":8080", r)

}
