/*
Package store
*/
package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/soulgeo/todolist/internal/todo"
)

const fileName string = "lists.json"

type JSONStore struct {
	File string
}

func NewJSONStore(file string) todo.Store {
	return &JSONStore{File: file}
}

func (jsonstore *JSONStore) Create(list todo.TodoList) error {
	lists, err := ParseJSON(fileName)
	if err != nil {
		return err
	}
	lists.Lists = append(lists.Lists, list)
	err = WriteJSON(fileName, *lists)
	if err != nil {
		return err
	}
	return nil
}

func (jsonstore *JSONStore) Delete(listname string) error {
	lists, err := ParseJSON(fileName)
	if err != nil {
		return err
	}
	index, _ := todo.FindList(listname, *lists)
	if index == -1 {
		return errors.New("list with given name not found")
	}
	lists.Lists = append(lists.Lists[:index], lists.Lists[index+1:]...)
	err = WriteJSON(fileName, *lists)
	if err != nil {
		return err
	}
	return nil
}

func (jsonstore *JSONStore) Add(listname string, item todo.Item) error {
	lists, err := ParseJSON(fileName)
	if err != nil {
		return err
	}
	index, list := todo.FindList(listname, *lists)
	if index == -1 {
		return errors.New("list with given name not found")
	}
	list.Items = append(list.Items, item)
	err = WriteJSON(fileName, *lists)
	if err != nil {
		return err
	}
	return nil
}

func (jsonstore *JSONStore) Remove(listname string, index int) error {
	lists, err := ParseJSON(fileName)
	if err != nil {
		return err
	}
	index, list := todo.FindList(listname, *lists)
	if index == -1 {
		return errors.New("list with given name not found")
	}
	list.Items = append(list.Items[:index], list.Items[index+1:]...)
	err = WriteJSON(fileName, *lists)
	if err != nil {
		return err
	}
	return nil
}

func (jsonstore *JSONStore) Get(listname string) (*todo.TodoList, error) {
	lists, err := ParseJSON(fileName)
	if err != nil {
		return nil, err
	}
	for _, l := range lists.Lists {
		if l.Name == listname {
			return &l, nil
		}
	}
	return nil, errors.New("list with given name not found")
}

func ParseJSON(filename string) (*todo.ListOfLists, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lists := todo.ListOfLists{}
	err = json.Unmarshal(content, &lists)
	if err != nil {
		return nil, err
	}
	return &lists, nil
}

func WriteJSON(filename string, lists todo.ListOfLists) error {
	in, err := json.Marshal(lists)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, in, 0o666)
	if err != nil {
		return err
	}
	return nil
}
