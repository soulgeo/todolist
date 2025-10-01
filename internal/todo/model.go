/*
Package todo
*/
package todo

import "time"

type Item struct {
	Title     string `json:"title"`
	Priority  string `json:"priority"`
	Completed bool   `json:"completed"`
}

type TodoList struct {
	Name         string    `json:"name"`
	Items        []Item    `json:"items"`
	CreationTime time.Time `json:"creationtime"`
	LastEdited   time.Time `json:"lastedited"`
}

type ListOfLists struct {
	Lists []TodoList `json:"lists"`
}
