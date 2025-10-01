package todo

type Store interface {
	Create(list TodoList) error
	Delete(listname string) error
	Add(listname string, item Item) error
	Remove(listname string, index int) error
	Get(listname string) (*TodoList, error)
}
