package list

func Wrap[T any](items []T) *Wrapper[T] {
	return &Wrapper[T]{
		items: items,
	}
}

type Wrapper[T any] struct {
	items []T
}

func (w *Wrapper[T]) Each(consumer func(T)) {
	Each(w.items,consumer)
}

func (w *Wrapper[T]) EachIndex(consumer func(int, T)) {
	EachIndex(w.items,consumer)
}