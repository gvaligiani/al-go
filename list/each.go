package list

func Each[T any](items []T, consumer func(T)) {
	for _, item := range items {
		consumer(item)
	}
}