package main

import (
	"github.com/soulgeo/todolist/internal/app" // ← adjust to your module path / package
	"github.com/soulgeo/todolist/internal/store"
	"github.com/soulgeo/todolist/internal/todo"
)

func main() {
	jsonstore := store.NewJSONStore("/path/to/todolist.json")
	app.SetService(todo.NewService(jsonstore))
	app.Execute()
}
