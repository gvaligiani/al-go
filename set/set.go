package set

// alias

type Set[T comparable] map[T]struct{}

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

func (s Set[T]) AllOf(predicate Predicate[T]) bool {
	return AllOf(s, predicate)
}

func (s Set[T]) AnyOf(predicate Predicate[T]) bool {
	return AnyOf(s, predicate)
}

func (s Set[T]) NoneOf(predicate Predicate[T]) bool {
	return NoneOf(s, predicate)
}

// each

func (s Set[T]) Each(consumer Consumer[T]) {
	Each(s, consumer)
}

// find

func (s Set[T]) Find(value T) bool {
	return Find(s, value)
}

func (s Set[T]) FindIf(predicate Predicate[T]) (T, bool) {
	return FindIf(s, predicate)
}

func (s Set[T]) FindIfNot(predicate Predicate[T]) (T, bool) {
	return FindIfNot(s, predicate)
}

// copy

func (s Set[T]) Copy() Set[T] {
	return Set[T](Copy(s))
}

func (s Set[T]) CopyIf(predicate Predicate[T]) Set[T] {
	return Set[T](CopyIf(s, predicate))
}

func (s Set[T]) CopyIfNot(predicate Predicate[T]) Set[T] {
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
	if s == nil || len(*s) == 0 {
		return false
	}
	if _, ok := (*s)[item]; ok {
		delete(*s, item)
		return true
	}
	return false
}

func (s *Set[T]) Clear() bool {
	if s == nil || len(*s) == 0 {
		return false
	}
	*s = Set[T]{}
	return true
}

func (s *Set[T]) RemoveIf(predicate Predicate[T]) bool {
	if s == nil || len(*s) == 0 {
		return false
	}
	return RemoveIf(s, predicate)
}

func (s *Set[T]) KeepIf(predicate Predicate[T]) bool {
	if s == nil || len(*s) == 0 {
		return false
	}
	return KeepIf(s, predicate)
}
