package set

// //////////////////////////////////////////////////
// has

func Has[T comparable](items map[T]struct{}, value T) bool {
	if items == nil {
		return false
	}
	_, ok := items[value]
	return ok
}
