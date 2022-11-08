package list

import "github.com/gvaligiani/al.go/util"

// alias

type List[V comparable] []V

// builder

func New[V comparable](values ...V) List[V] {
	// TODO add capacity
	// TODO remove values
	l := make(List[V], 0, len(values))
	l = append(l, values...)
	return l
}

func (l List[V]) With(value V) List[V] {
	l.Add(value)
	return l
}

// getter

// state

func (l List[V]) IsEmpty() bool {
	return len(l) == 0
}

func (l List[V]) AllOf(predicate util.Predicate[V]) bool {
	return AllOf(l, predicate)
}

func (l List[V]) AllIndexOf(predicate util.BiPredicate[int, V]) bool {
	return AllIndexOf(l, predicate)
}

func (l List[V]) AnyOf(predicate util.Predicate[V]) bool {
	return AnyOf(l, predicate)
}

func (l List[V]) AnyIndexOf(predicate util.BiPredicate[int, V]) bool {
	return AnyIndexOf(l, predicate)
}

func (l List[V]) NoneOf(predicate util.Predicate[V]) bool {
	return NoneOf(l, predicate)
}

func (l List[V]) NoIndexOf(predicate util.BiPredicate[int, V]) bool {
	return NoIndexOf(l, predicate)
}

// each

func (l List[V]) Each(consumer util.Consumer[V]) {
	Each(l, consumer)
}

func (l List[V]) EachIndex(consumer util.BiConsumer[int, V]) {
	EachIndex(l, consumer)
}

// find

func (l List[V]) FindIndex(index int) bool {
	return FindIndex(l, index)
}

func (l List[V]) FindValueFromIndex(index int) (V, bool) {
	return FindValueFromIndex(l, index)
}

func (l List[V]) Find(value V) bool {
	return Find(l, value)
}

func (l List[V]) FindIndexFromValue(value V) (int, bool) {
	return FindIndexFromValue(l, value)
}

func (l List[V]) FindIf(predicate util.Predicate[V]) (V, bool) {
	return FindIf(l, predicate)
}

func (l List[V]) FindIfIndex(predicate util.BiPredicate[int, V]) (int, V, bool) {
	return FindIfIndex(l, predicate)
}

func (l List[V]) FindIfNot(predicate util.Predicate[V]) (V, bool) {
	return FindIfNot(l, predicate)
}

func (l List[V]) FindIfNotIndex(predicate util.BiPredicate[int, V]) (int, V, bool) {
	return FindIfNotIndex(l, predicate)
}

// copy

func (l List[V]) Copy() List[V] {
	return List[V](Copy(l))
}

func (l List[V]) CopyIf(predicate util.Predicate[V]) List[V] {
	return List[V](CopyIf(l, predicate))
}

func (l List[V]) CopyIfIndex(predicate util.BiPredicate[int, V]) List[V] {
	return List[V](CopyIfIndex(l, predicate))
}

func (l List[V]) CopyIfNot(predicate util.Predicate[V]) List[V] {
	return List[V](CopyIfNot(l, predicate))
}

func (l List[V]) CopyIfNotIndex(predicate util.BiPredicate[int, V]) List[V] {
	return List[V](CopyIfNotIndex(l, predicate))
}

// modifier

func (l *List[V]) Add(value V) bool {
	*l = append(*l, value)
	return true
}

func (l *List[V]) Remove(value V) bool {
	return RemoveIf(l, func(v V) bool { return util.Equal(v, value) })
}

func (l *List[V]) Clear() bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	*l = List[V]{}
	return true
}

func (l *List[V]) RemoveIf(predicate util.Predicate[V]) bool {
	return RemoveIf(l, predicate)
}

func (l *List[V]) RemoveIfIndex(predicate util.BiPredicate[int, V]) bool {
	return RemoveIfIndex(l, predicate)
}

func (l *List[V]) KeepIf(predicate util.Predicate[V]) bool {
	return KeepIf(l, predicate)
}

func (l *List[V]) KeepIfIndex(predicate util.BiPredicate[int, V]) bool {
	return KeepIfIndex(l, predicate)
}
