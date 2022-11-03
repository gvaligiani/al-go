package list

func RemoveIf[T any, L ~[]T](items *L, predicate Predicate[T]) {
	if len(*items) == 0 {
		return
	}
	var empty T
	size := len(*items)
	nbRemoved := 0
	for i := 0; i < size-nbRemoved; {
		if predicate((*items)[i]) {
			nbRemoved++
			(*items)[i] = (*items)[size-nbRemoved]
			(*items)[size-nbRemoved] = empty
		} else {
			i++
		}
	}
	(*items) = (*items)[:size-nbRemoved]
}
