package list

import "github.com/gvaligiani/al.go/util"

// alias

type DeepList[T any] []T

// builder

func NewDeep[T any](values ...T) DeepList[T] {
	l := make(DeepList[T], 0, len(values))
	l = append(l, values...)
	return l
}

// getter

// state

func (l DeepList[T]) IsEmpty() bool {
	return len(l) == 0
}

func (l DeepList[T]) AllOf(predicate util.Predicate[T]) bool {
	return AllOf(l, predicate)
}

func (l DeepList[T]) AllIndexOf(predicate util.BiPredicate[int, T]) bool {
	return AllIndexOf(l, predicate)
}

func (l DeepList[T]) AnyOf(predicate util.Predicate[T]) bool {
	return AnyOf(l, predicate)
}

func (l DeepList[T]) AnyIndexOf(predicate util.BiPredicate[int, T]) bool {
	return AnyIndexOf(l, predicate)
}

func (l DeepList[T]) NoneOf(predicate util.Predicate[T]) bool {
	return NoneOf(l, predicate)
}

func (l DeepList[T]) NoIndexOf(predicate util.BiPredicate[int, T]) bool {
	return NoIndexOf(l, predicate)
}

// each

func (l DeepList[T]) Each(consumer util.Consumer[T]) {
	Each(l, consumer)
}

func (l DeepList[T]) EachIndex(consumer util.BiConsumer[int, T]) {
	EachIndex(l, consumer)
}

// find

func (l DeepList[T]) Find(value T) bool {
	return DeepFind(l, value)
}

func (l DeepList[T]) FindIndex(value T) (int, bool) {
	return DeepFindIndexOf(l, value)
}

func (l DeepList[T]) FindIf(predicate util.Predicate[T]) (T, bool) {
	return FindIf(l, predicate)
}

func (l DeepList[T]) FindIndexIf(predicate util.BiPredicate[int, T]) (int, T, bool) {
	return FindIndexIf(l, predicate)
}

func (l DeepList[T]) FindIfNot(predicate util.Predicate[T]) (T, bool) {
	return FindIfNot(l, predicate)
}

func (l DeepList[T]) FindIndexIfNot(predicate util.BiPredicate[int, T]) (int, T, bool) {
	return FindIndexIfNot(l, predicate)
}

// copy

func (l DeepList[T]) Copy() DeepList[T] {
	return DeepList[T](Copy(l))
}

func (l DeepList[T]) CopyIf(predicate util.Predicate[T]) DeepList[T] {
	return DeepList[T](CopyIf(l, predicate))
}

func (l DeepList[T]) CopyIfNot(predicate util.Predicate[T]) DeepList[T] {
	return DeepList[T](CopyIfNot(l, predicate))
}

// modifier

func (l *DeepList[T]) Add(value T) bool {
	*l = append(*l, value)
	return true
}

func (l *DeepList[T]) Remove(value T) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return RemoveIf(l, func(t T) bool { return util.DeepEqual(t, value) })
}

func (l *DeepList[T]) Clear() bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	*l = DeepList[T]{}
	return true
}

func (l *DeepList[T]) RemoveIf(predicate util.Predicate[T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return RemoveIf(l, predicate)
}

func (l *DeepList[T]) RemoveIndexIf(predicate util.BiPredicate[int, T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return RemoveIndexIf(l, predicate)
}

func (l *DeepList[T]) KeepIf(predicate util.Predicate[T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return KeepIf(l, predicate)
}

func (l *DeepList[T]) KeepIndexIf(predicate util.BiPredicate[int, T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return KeepIndexIf(l, predicate)
}
