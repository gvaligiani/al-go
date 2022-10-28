package list

func EachIndex[T any](items []T, consumer func(int, T)) {
	for index, item := range items {
		consumer(index, item)
	}
}