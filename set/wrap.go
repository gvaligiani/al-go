package set

func Wrap[T comparable](items map[T]struct{}) *Wrapper[T] {
	return &Wrapper[T]{
		items: items,
	}
}

type Wrapper[T comparable] struct {
	items map[T]struct{}
}

func (w *Wrapper[T]) Add(item T) {
	w.items[item] = struct{}{}
}

func (w *Wrapper[T]) AllOf(predicate func(T) bool) bool {
	return AllOf[T](w.items,predicate)
}

func (w *Wrapper[T]) Each(consumer func(T)) {
	Each[T](w.items,consumer)
}