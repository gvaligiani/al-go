package list

func Each[T any, L ~[]T](items L, consumer Consumer[T]) {
	for _, item := range items {
		consumer(item)
	}
}
