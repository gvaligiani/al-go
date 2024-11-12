package list

// //////////////////////////////////////////////////
// wrap

func Wrap[T comparable](items []T) List[T] {
	return List[T](items)
}
