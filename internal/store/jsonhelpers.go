package store

import (
	"encoding/json"
	"os"

	"github.com/soulgeo/todolist/internal/todo"
)

func ParseJSON(filename string) (*todo.ListOfLists, error) {
	out, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &todo.ListOfLists{}, nil
		}
		return nil, err
	}
	lists := todo.ListOfLists{}
	err = json.Unmarshal(out, &lists)
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
