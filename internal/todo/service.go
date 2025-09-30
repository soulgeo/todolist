package todo

type Service struct {
	store Store
}

func NewService(s Store) *Service {
	return &Service{store: s}
}

func (svc *Service) Create(list TodoList) error {
	return svc.store.Create(list)
}

func (svc *Service) Delete(name string) error {
	return svc.store.Delete(name)
}

func (svc *Service) Add(list TodoList, item Item) error {
	return svc.store.Add(list, item)
}

func (svc *Service) Remove(list TodoList, index int) error {
	return svc.store.Remove(list, index)
}
