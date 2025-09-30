package todo

type Store interface {
	Create(list TodoList) error
	Delete(name string) error
	Add(list TodoList, item Item) error
	Remove(list TodoList, index int) error
}
