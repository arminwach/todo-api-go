package models

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var Todos = []Todo{
	{ID: 1, Task: "Clean Code"},
	{ID: 2, Task: "Eat"},
}
