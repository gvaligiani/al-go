package set

func Wrap[T comparable](items map[T]struct{}) *Wrapper[T] {
	return &Wrapper[T]{
		items: items,
	}
}

type Wrapper[T comparable] struct {
	items map[T]struct{}
}

func (w *Wrapper[T]) Each(consumer func(T)) {
	Each(w.items,consumer)
}