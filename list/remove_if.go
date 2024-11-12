package list

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// remove if

func RemoveIf[T any, L ~[]T](items *L, predicate fn.Predicate[T]) {
	if items == nil {
		return
	}
	size := len(*items)
	if size == 0 {
		return
	}
	// note: this method does not keep the original order of the list
	var empty T
	indexEnd := size - 1
	for indexStart := 0; indexStart <= indexEnd; {
		if predicate((*items)[indexStart]) {
			// move last element at current index
			(*items)[indexStart] = (*items)[indexEnd]
			(*items)[indexEnd] = empty
			indexEnd--
		} else {
			indexStart++
		}
	}
	// reduce the list to the right size
	(*items) = (*items)[:indexEnd+1]
}
