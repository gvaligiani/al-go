package dict

// //////////////////////////////////////////////////
// new

func New[K comparable, V any](items ...pair[K, V]) Dict[K, V] {
	dict := make(Dict[K, V], 0)
	for _, item := range items {
		dict[item.Key] = item.Value
	}
	return dict
}
