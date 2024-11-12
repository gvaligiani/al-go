package dict

// //////////////////////////////////////////////////
// wrap

func Wrap[K comparable, V any](items map[K]V) Dict[K, V] {
	return Dict[K, V](items)
}
