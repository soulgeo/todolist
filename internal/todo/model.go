/*
Package todo
*/
package todo

type Item struct {
	Title    string `json:"title"`
	Priority string `json:"priority"`
}

type TodoList struct {
	Name  string `json:"name"`
	Items []Item `json:"items"`
}
