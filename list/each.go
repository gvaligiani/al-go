package list

func Each[T any, L ~[]T](items L, consumer Consumer[T]) {
	for index, item := range items {
		consumer(index, item)
	}
}
