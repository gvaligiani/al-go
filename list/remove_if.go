package list

import "github.com/gvaligiani/al.go/util"

func RemoveIf[V any, L ~[]V](l *L, predicate util.Predicate[V]) bool {
	return RemoveIfIndex(l, util.TestOnSecondArg[int](predicate))
}

func RemoveIfIndex[V any, L ~[]V](l *L, predicate util.BiPredicate[int, V]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	// note:
	//  - this method does not keep the original order of the list
	//  - but it sends the right old index to the predicate
	var empty V
	size := len(*l)
	index := 0
	indexEnd := size - 1
	for indexStart := 0; indexStart <= indexEnd; {
		if predicate(index, (*l)[indexStart]) {
			// move last element at current index
			(*l)[indexStart] = (*l)[indexEnd]
			(*l)[indexEnd] = empty
			index = indexEnd
			indexEnd--
		} else {
			indexStart++
			index = indexStart
		}
	}
	// reduce the list to the right size
	(*l) = (*l)[:indexEnd+1]
	return indexEnd < size-1
}
