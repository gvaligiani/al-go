package set

// alias

type Set[T comparable] map[T]struct{}
type Predicate[T comparable] func(T) bool
type Consumer[T comparable] func(T)
type Supplier[T comparable] func() T
type Transformer[T comparable, O any] func(T) O

// predicate

func Not[T comparable](predicate Predicate[T]) Predicate[T] {
	return func(item T) bool {
		return !predicate(item)
	}
}

func True[T comparable]() Predicate[T] {
	return func(_ T) bool {
		return true
	}
}

// builder

func New[T comparable](items ...T) Set[T] {
	s := Set[T]{}
	for _, item := range items {
		s.Add(item)
	}
	return s
}

// getter

func (s Set[T]) Values() []T {
	l := make([]T, 0, len(s))
	for item := range s {
		l = append(l, item)
	}
	return l
}

// state

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) AllOf(predicate func(T) bool) bool {
	return AllOf(s, predicate)
}

func (s Set[T]) AnyOf(predicate func(T) bool) bool {
	return AnyOf(s, predicate)
}

func (s Set[T]) NoneOf(predicate func(T) bool) bool {
	return NoneOf(s, predicate)
}

// each

func (s Set[T]) Each(consumer func(T)) {
	Each(s, consumer)
}

// find

func (s Set[T]) Find(value T) bool {
	return Find(s, value)
}

func (s Set[T]) FindIf(predicate func(T) bool) (T, bool) {
	return FindIf(s, predicate)
}

func (s Set[T]) FindIfNot(predicate func(T) bool) (T, bool) {
	return FindIfNot(s, predicate)
}

// copy

func (s Set[T]) Copy() Set[T] {
	return Set[T](Copy(s))
}

func (s Set[T]) CopyIf(predicate func(T) bool) Set[T] {
	return Set[T](CopyIf(s, predicate))
}

func (s Set[T]) CopyIfNot(predicate func(T) bool) Set[T] {
	return Set[T](CopyIfNot(s, predicate))
}

// modifier

func (s *Set[T]) Add(item T) bool {
	if s == nil {
		return false
	}
	if _, ok := (*s)[item]; !ok {
		(*s)[item] = struct{}{}
		return true
	}
	return false
}

func (s *Set[T]) Remove(item T) bool {
	if s == nil {
		return false
	}
	if _, ok := (*s)[item]; ok {
		delete(*s, item)
		return true
	}
	return false
}

func (s *Set[T]) Clear() {
	*s = Set[T]{}
}

func (s Set[T]) RemoveIf(predicate func(T) bool) {
	RemoveIf(&s, predicate)
}

func (s Set[T]) KeepIf(predicate func(T) bool) {
	KeepIf(&s, predicate)
}
