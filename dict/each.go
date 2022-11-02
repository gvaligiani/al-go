package dict

func Each[K comparable, T any, M ~map[K]T](items M, consumer Consumer[K, T]) {
	for key, item := range items {
		consumer(key, item)
	}
}
