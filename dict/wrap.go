package dict

func Wrap[K comparable,T any](items map[K]T) *Wrapper[K,T] {
	return &Wrapper[K,T]{
		items: items,
	}
}

type Wrapper[K comparable,T any] struct {
	items map[K]T
}

func (w *Wrapper[K,T]) Each(consumer func(K, T)) {
	Each(w.items,consumer)
}