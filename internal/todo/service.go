package todo

import "time"

type Service struct {
	store Store
}

func NewService(s Store) *Service {
	return &Service{store: s}
}

func (svc *Service) Create(name string) error {
	newList := TodoList{
		Name:         name,
		Items:        make([]Item, 0),
		CreationTime: time.Now(),
		LastEdited:   time.Now(),
	}
	return svc.store.Create(newList)
}

func (svc *Service) Delete(name string) error {
	return svc.store.Delete(name)
}

func (svc *Service) Add(
	listname string,
	itemname string,
	priority string,
) error {
	item := Item{
		Title:     itemname,
		Priority:  priority,
		Completed: false,
	}
	return svc.store.Add(listname, item)
}

func (svc *Service) Remove(list TodoList, index int) error {
	return svc.store.Remove(list, index)
}
