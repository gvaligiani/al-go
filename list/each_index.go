package list

func EachIndex[T any, L ~[]T](items L, consumer func(int, T)) {
	for index, item := range items {
		consumer(index, item)
	}
}
