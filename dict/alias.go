package dict

// alias

type Dict[K comparable, T any] map[K]T
type Predicate[K comparable, T any] func(K, T) bool
type Consumer[K comparable, T any] func(K, T)
type Supplier[K comparable, T any] func() (K, T)
type Transformer[K comparable, T any, O any] func(K, T) O

// predicate

func Not[K comparable, T any](predicate Predicate[K, T]) Predicate[K, T] {
	return func(key K, item T) bool {
		return !predicate(key, item)
	}
}

func True[K comparable, T any]() Predicate[K, T] {
	return func(_ K, _ T) bool {
		return true
	}
}

// builder

func New[K comparable, T any]() Dict[K, T] {
	return Dict[K, T]{}
}

// getter

func (d Dict[K, T]) Keys() []K {
	l := make([]K, 0, len(d))
	for key := range d {
		l = append(l, key)
	}
	return l
}

func (d Dict[K, T]) Values() []T {
	l := make([]T, 0, len(d))
	for _, item := range d {
		l = append(l, item)
	}
	return l
}

// state

func (d Dict[K, T]) IsEmpty() bool {
	return len(d) == 0
}

func (d Dict[K, T]) AllOf(predicate Predicate[K, T]) bool {
	return AllOf(d, predicate)
}

func (d Dict[K, T]) AnyOf(predicate Predicate[K, T]) bool {
	return AnyOf(d, predicate)
}

func (d Dict[K, T]) NoneOf(predicate Predicate[K, T]) bool {
	return NoneOf(d, predicate)
}

// each

func (d Dict[K, T]) Each(consumer func(K, T)) {
	Each(d, consumer)
}

// find

func (d Dict[K, T]) FindKey(key K) bool {
	return FindKey(d, key)
}

func (d Dict[K, T]) FindValue(value T) bool {
	return FindValue(d, value)
}

func (d Dict[K, T]) FindIf(predicate Predicate[K, T]) (K, T, bool) {
	return FindIf(d, predicate)
}

func (d Dict[K, T]) FindIfNot(predicate Predicate[K, T]) (K, T, bool) {
	return FindIfNot(d, predicate)
}

// copy

func (d Dict[K, T]) Copy() Dict[K, T] {
	return Dict[K, T](Copy(d))
}

func (d Dict[K, T]) CopyIf(predicate Predicate[K, T]) Dict[K, T] {
	return Dict[K, T](CopyIf(d, predicate))
}

func (d Dict[K, T]) CopyIfNot(predicate Predicate[K, T]) Dict[K, T] {
	return Dict[K, T](CopyIfNot(d, predicate))
}

// modifier

func (d *Dict[K, T]) Add(key K, item T) bool {
	_, overriden := (*d)[key]
	(*d)[key] = item
	return overriden
}

func (d *Dict[K, T]) Remove(key K) bool {
	if _, ok := (*d)[key]; ok {
		delete(*d, key)
		return true
	}
	return false
}

func (d *Dict[K, T]) Clear() {
	*d = Dict[K, T]{}
}

func (d *Dict[K, T]) RemoveIf(predicate Predicate[K, T]) {
	RemoveIf(d, predicate)
}

func (d *Dict[K, T]) KeepIf(predicate Predicate[K, T]) {
	KeepIf(d, predicate)
}
