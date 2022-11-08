package list

func RemoveIf[T any, L ~[]T](items *L, predicate Predicate[T]) bool {
	if len(*items) == 0 {
		return false
	}
	// note:
	//  - this method does not keep the original order of the list
	//  - but it sends the right old index to the predicate
	var empty T
	size := len(*items)
	index := 0
	indexEnd := size - 1
	for indexStart := 0; indexStart <= indexEnd; {
		if predicate(index, (*items)[indexStart]) {
			// move last element at current index
			(*items)[indexStart] = (*items)[indexEnd]
			(*items)[indexEnd] = empty
			index = indexEnd
			indexEnd--
		} else {
			indexStart++
			index = indexStart
		}
	}
	// reduce the list to the right size
	(*items) = (*items)[:indexEnd+1]
	return indexEnd < size-1
}
