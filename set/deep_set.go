package set

import "github.com/gvaligiani/al.go/util"

// alias

type DeepSet[V comparable] map[V]struct{}

// builder

func NewDeep[V comparable](values ...V) DeepSet[V] {
	// TODO add capacity
	// TODO remove values
	s := DeepSet[V]{}
	for _, v := range values {
		s.Add(v)
	}
	return s
}

func (s DeepSet[V]) With(value V) DeepSet[V] {
	s.Add(value)
	return s
}

// getter

func (s DeepSet[V]) Values() []V {
	l := make([]V, 0, len(s))
	for v := range s {
		l = append(l, v)
	}
	return l
}

// state

func (s DeepSet[V]) IsEmpty() bool {
	return len(s) == 0
}

func (s DeepSet[V]) AllOf(predicate util.Predicate[V]) bool {
	return AllOf(s, predicate)
}

func (s DeepSet[V]) AnyOf(predicate util.Predicate[V]) bool {
	return AnyOf(s, predicate)
}

func (s DeepSet[V]) NoneOf(predicate util.Predicate[V]) bool {
	return NoneOf(s, predicate)
}

// each

func (s DeepSet[V]) Each(consumer util.Consumer[V]) {
	Each(s, consumer)
}

// find

func (s DeepSet[V]) Find(value V) bool {
	return DeepFind(s, value)
}

func (s DeepSet[V]) FindIf(predicate util.Predicate[V]) (V, bool) {
	return FindIf(s, predicate)
}

func (s DeepSet[V]) FindIfNot(predicate util.Predicate[V]) (V, bool) {
	return FindIfNot(s, predicate)
}

// copy

func (s DeepSet[V]) Copy() DeepSet[V] {
	return Copy(s)
}

func (s DeepSet[V]) CopyIf(predicate util.Predicate[V]) DeepSet[V] {
	return CopyIf(s, predicate)
}

func (s DeepSet[V]) CopyIfNot(predicate util.Predicate[V]) DeepSet[V] {
	return CopyIfNot(s, predicate)
}

// modifier

func (s *DeepSet[V]) Add(value V) bool {
	if s == nil {
		return false
	}
	if !DeepFind(*s, value) {
		(*s)[value] = struct{}{}
		return true
	}
	return false
}

func (s *DeepSet[V]) Remove(value V) bool {
	if s == nil || len(*s) == 0 {
		return false
	}
	return RemoveIf(s, func(v V) bool { return util.DeepEqual(v, value) })
}

func (s *DeepSet[V]) Clear() bool {
	if s == nil || len(*s) == 0 {
		return false
	}
	*s = DeepSet[V]{}
	return true
}

func (s *DeepSet[V]) RemoveIf(predicate util.Predicate[V]) bool {
	return RemoveIf(s, predicate)
}

func (s *DeepSet[V]) KeepIf(predicate util.Predicate[V]) bool {
	return KeepIf(s, predicate)
}

func (s DeepSet[V]) Union(right DeepSet[V]) DeepSet[V] {
	return DeepUnion(s, right)
}

func (s DeepSet[V]) Difference(right DeepSet[V]) DeepSet[V] {
	return DeepDifference(s, right)
}

func (s DeepSet[V]) Intersection(right DeepSet[V]) DeepSet[V] {
	return DeepIntersection(s, right)
}
