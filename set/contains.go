package set

// //////////////////////////////////////////////////
// contains

func Contains[T comparable](items map[T]struct{}, value T) bool {
	if items == nil {
		return false
	}
	_, ok := items[value]
	return ok
}
