package todo

import (
	"time"

	"github.com/soulgeo/todolist/internal/config"
)

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

func (svc *Service) Remove(listname string, index int) error {
	return svc.store.Remove(listname, index)
}

func (svc *Service) GetList(listname string) (*TodoList, error) {
	return svc.store.GetList(listname)
}

func (svc *Service) GetSelectedList() (*TodoList, error) {
	list, err := config.GetSelected()
	if err != nil {
		return nil, err
	}
	return svc.store.GetList(list)
}

func (svc *Service) Select(listname string) error {
	return config.SetSelected(listname)
}
