package todo

func FindList(listname string, lists ListOfLists) (int, *TodoList) {
	for i := range lists.Lists {
		if lists.Lists[i].Name == listname {
			return i, &lists.Lists[i]
		}
	}
	return -1, nil
}
