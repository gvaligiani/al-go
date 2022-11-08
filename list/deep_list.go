package list

import "github.com/gvaligiani/al.go/util"

// alias

type DeepList[V any] []V

// builder

func NewDeep[V any](values ...V) DeepList[V] {
	// TODO remove values
	// TODO add capacity
	l := make(DeepList[V], 0, len(values))
	l = append(l, values...)
	return l
}

func (l DeepList[V]) With(value V) DeepList[V] {
	l.Add(value)
	return l
}

// getter

// state

func (l DeepList[V]) IsEmpty() bool {
	return len(l) == 0
}

func (l DeepList[V]) AllOf(predicate util.Predicate[V]) bool {
	return AllOf(l, predicate)
}

func (l DeepList[V]) AllIndexOf(predicate util.BiPredicate[int, V]) bool {
	return AllIndexOf(l, predicate)
}

func (l DeepList[V]) AnyOf(predicate util.Predicate[V]) bool {
	return AnyOf(l, predicate)
}

func (l DeepList[V]) AnyIndexOf(predicate util.BiPredicate[int, V]) bool {
	return AnyIndexOf(l, predicate)
}

func (l DeepList[V]) NoneOf(predicate util.Predicate[V]) bool {
	return NoneOf(l, predicate)
}

func (l DeepList[V]) NoIndexOf(predicate util.BiPredicate[int, V]) bool {
	return NoIndexOf(l, predicate)
}

// each

func (l DeepList[V]) Each(consumer util.Consumer[V]) {
	Each(l, consumer)
}

func (l DeepList[V]) EachIndex(consumer util.BiConsumer[int, V]) {
	EachIndex(l, consumer)
}

// find

func (l DeepList[V]) FindIndex(index int) bool {
	return FindIndex(l, index)
}

func (l DeepList[V]) FindValueFromIndex(index int) (V, bool) {
	return FindValueFromIndex(l, index)
}

func (l DeepList[V]) Find(value V) bool {
	return DeepFind(l, value)
}

func (l DeepList[V]) FindIndexFromValue(value V) (int, bool) {
	return DeepFindIndexFromValue(l, value)
}

func (l DeepList[V]) FindIf(predicate util.Predicate[V]) (V, bool) {
	return FindIf(l, predicate)
}

func (l DeepList[V]) FindIfIndex(predicate util.BiPredicate[int, V]) (int, V, bool) {
	return FindIfIndex(l, predicate)
}

func (l DeepList[V]) FindIfNot(predicate util.Predicate[V]) (V, bool) {
	return FindIfNot(l, predicate)
}

func (l DeepList[V]) FindIfNotIndex(predicate util.BiPredicate[int, V]) (int, V, bool) {
	return FindIfNotIndex(l, predicate)
}

// copy

func (l DeepList[V]) Copy() DeepList[V] {
	return DeepList[V](Copy(l))
}

func (l DeepList[V]) CopyIf(predicate util.Predicate[V]) DeepList[V] {
	return DeepList[V](CopyIf(l, predicate))
}

func (l DeepList[V]) CopyIfIndex(predicate util.BiPredicate[int, V]) DeepList[V] {
	return DeepList[V](CopyIfIndex(l, predicate))
}

func (l DeepList[V]) CopyIfNot(predicate util.Predicate[V]) DeepList[V] {
	return DeepList[V](CopyIfNot(l, predicate))
}

func (l DeepList[V]) CopyIfNotIndex(predicate util.BiPredicate[int, V]) DeepList[V] {
	return DeepList[V](CopyIfNotIndex(l, predicate))
}

// modifier

func (l *DeepList[V]) Add(value V) bool {
	// TODO append multi value
	*l = append(*l, value)
	return true
}

func (l *DeepList[V]) Remove(value V) bool {
	return RemoveIf(l, func(v V) bool { return util.DeepEqual(v, value) })
}

func (l *DeepList[V]) Clear() bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	*l = DeepList[V]{}
	return true
}

func (l *DeepList[V]) RemoveIf(predicate util.Predicate[V]) bool {
	return RemoveIf(l, predicate)
}

func (l *DeepList[V]) RemoveIfIndex(predicate util.BiPredicate[int, V]) bool {
	return RemoveIfIndex(l, predicate)
}

func (l *DeepList[V]) KeepIf(predicate util.Predicate[V]) bool {
	return KeepIf(l, predicate)
}

func (l *DeepList[V]) KeepIfIndex(predicate util.BiPredicate[int, V]) bool {
	return KeepIfIndex(l, predicate)
}
