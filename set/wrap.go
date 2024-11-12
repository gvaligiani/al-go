package set

// //////////////////////////////////////////////////
// wrap

func Wrap[T comparable](items map[T]struct{}) Set[T] {
	return Set[T](items)
}
