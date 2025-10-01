package todo

func FindList(listname string, lists ListOfLists) (int, *TodoList) {
	for i, l := range lists.Lists {
		if l.Name == listname {
			return i, &l
		}
	}
	return -1, nil
}
