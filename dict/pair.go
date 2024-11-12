package dict

// //////////////////////////////////////////////////
// pair

type pair[K comparable, V any] struct {
	Key   K
	Value V
}

func Pair[K comparable, V any](key K, value V) pair[K, V] {
	return pair[K, V]{
		Key:   key,
		Value: value,
	}
}
