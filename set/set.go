package set

import "github.com/gvaligiani/al.go/util"

// alias

type Set[V comparable] map[V]struct{}

// TODO implement union
// TODO implement diff
// TODO implement intersection

// builder

func New[V comparable](values ...V) Set[V] {
	// TODO add capacity
	// TODO remove values
	s := Set[V]{}
	for _, v := range values {
		s.Add(v)
	}
	return s
}

func (s Set[V]) With(value V) Set[V] {
	s.Add(value)
	return s
}

// getter

func (s Set[V]) Values() []V {
	l := make([]V, 0, len(s))
	for v := range s {
		l = append(l, v)
	}
	return l
}

// state

func (s Set[V]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[V]) AllOf(predicate util.Predicate[V]) bool {
	return AllOf(s, predicate)
}

func (s Set[V]) AnyOf(predicate util.Predicate[V]) bool {
	return AnyOf(s, predicate)
}

func (s Set[V]) NoneOf(predicate util.Predicate[V]) bool {
	return NoneOf(s, predicate)
}

// each

func (s Set[V]) Each(consumer util.Consumer[V]) {
	Each(s, consumer)
}

// find

func (s Set[V]) Find(value V) bool {
	return Find(s, value)
}

func (s Set[V]) FindIf(predicate util.Predicate[V]) (V, bool) {
	return FindIf(s, predicate)
}

func (s Set[V]) FindIfNot(predicate util.Predicate[V]) (V, bool) {
	return FindIfNot(s, predicate)
}

// copy

func (s Set[V]) Copy() Set[V] {
	return Set[V](Copy(s))
}

func (s Set[V]) CopyIf(predicate util.Predicate[V]) Set[V] {
	return Set[V](CopyIf(s, predicate))
}

func (s Set[V]) CopyIfNot(predicate util.Predicate[V]) Set[V] {
	return Set[V](CopyIfNot(s, predicate))
}

// modifier

func (s *Set[V]) Add(value V) bool {
	if s == nil {
		return false
	}
	if _, ok := (*s)[value]; !ok {
		(*s)[value] = struct{}{}
		return true
	}
	return false
}

func (s *Set[V]) Remove(value V) bool {
	if s == nil || len(*s) == 0 {
		return false
	}
	if _, ok := (*s)[value]; ok {
		delete(*s, value)
		return true
	}
	return false
}

func (s *Set[V]) Clear() bool {
	if s == nil || len(*s) == 0 {
		return false
	}
	*s = Set[V]{}
	return true
}

func (s *Set[V]) RemoveIf(predicate util.Predicate[V]) bool {
	return RemoveIf(s, predicate)
}

func (s *Set[V]) KeepIf(predicate util.Predicate[V]) bool {
	return KeepIf(s, predicate)
}

func (s Set[V]) Union(right Set[V]) Set[V] {
	return Union(s, right)
}

func (s Set[V]) Difference(right Set[V]) Set[V] {
	return Difference(s, right)
}

func (s Set[V]) Intersection(right Set[V]) Set[V] {
	return Intersection(s, right)
}
