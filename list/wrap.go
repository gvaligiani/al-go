package list

func Wrap[T any](items []T) *Wrapper[T] {
	return &Wrapper[T]{
		items: items,
	}
}

type Wrapper[T any] struct {
	items []T
}

func (w *Wrapper[T]) Each(consumer func(int, T)) {
	Each(w.items,consumer)
}