/*
Package store
*/
package store

import (
	"errors"
	"time"

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
	list.LastEdited = time.Now()
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
	ind, list := todo.FindList(listname, *lists)
	if ind == -1 {
		return errors.New("list with given name not found")
	}
	list.Items = append(list.Items[:index], list.Items[index+1:]...)
	list.LastEdited = time.Now()
	err = WriteJSON(fileName, *lists)
	if err != nil {
		return err
	}
	return nil
}

func (jsonstore *JSONStore) GetList(listname string) (*todo.TodoList, error) {
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
