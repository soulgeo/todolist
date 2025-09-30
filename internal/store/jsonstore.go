/*
Package store
*/
package store

import "github.com/soulgeo/todolist/internal/todo"

type JSONStore struct {
	File string
}

func NewJSONStore(file string) todo.Store {
	return &JSONStore{File: file}
}

func (jsonstore *JSONStore) Create(list todo.TodoList) error {
	// Do Stuff
	return nil
}

func (jsonstore *JSONStore) Delete(name string) error {
	// Do Stuff
	return nil
}

func (jsonstore *JSONStore) Add(list todo.TodoList, item todo.Item) error {
	// Do Stuff
	return nil
}

func (jsonstore *JSONStore) Remove(list todo.TodoList, index int) error {
	// Do Stuff
	return nil
}
