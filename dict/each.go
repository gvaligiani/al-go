package dict

func Each[K comparable,T any](items map[K]T, consumer func(K, T)) {
	for key, item := range items {
		consumer(key, item)
	}
}