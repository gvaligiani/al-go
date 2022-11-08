package list

import (
	"github.com/gvaligiani/al.go/util"
)

// alias

type List[T comparable] []T

// builder

func New[T comparable](values ...T) List[T] {
	l := make(List[T], 0, len(values))
	l = append(l, values...)
	return l
}

// getter

// state

func (l List[T]) IsEmpty() bool {
	return len(l) == 0
}

func (l List[T]) AllOf(predicate util.Predicate[T]) bool {
	return AllOf(l, predicate)
}

func (l List[T]) AnyOf(predicate util.Predicate[T]) bool {
	return AnyOf(l, predicate)
}

func (l List[T]) NoneOf(predicate util.Predicate[T]) bool {
	return NoneOf(l, predicate)
}

// each

func (l List[T]) Each(consumer util.Consumer[T]) {
	Each(l, consumer)
}

// find

func (l List[T]) Find(value T) bool {
	return Find(l, value)
}

func (l List[T]) FindIndexOf(value T) (int, bool) {
	return FindIndexOf(l, value)
}

func (l List[T]) FindIf(predicate util.Predicate[T]) (T, bool) {
	return FindIf(l, predicate)
}

func (l List[T]) FindIndexIf(predicate util.BiPredicate[int, T]) (int, T, bool) {
	return FindIndexIf(l, predicate)
}

func (l List[T]) FindIfNot(predicate util.Predicate[T]) (T, bool) {
	return FindIfNot(l, predicate)
}

func (l List[T]) FindIndexIfNot(predicate util.BiPredicate[int, T]) (int, T, bool) {
	return FindIndexIfNot(l, predicate)
}

// copy

func (l List[T]) Copy() List[T] {
	return List[T](Copy(l))
}

func (l List[T]) CopyIf(predicate util.Predicate[T]) List[T] {
	return List[T](CopyIf(l, predicate))
}

func (l List[T]) CopyIfNot(predicate util.Predicate[T]) List[T] {
	return List[T](CopyIfNot(l, predicate))
}

// modifier

func (l *List[T]) Add(value T) bool {
	*l = append(*l, value)
	return true
}

func (l *List[T]) Remove(value T) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return RemoveIf(l, func(t T) bool { return util.Equal(t, value) })
}

func (l *List[T]) Clear() bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	*l = List[T]{}
	return true
}

func (l *List[T]) RemoveIf(predicate util.Predicate[T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return RemoveIf(l, predicate)
}

func (l *List[T]) RemoveIndexIf(predicate util.BiPredicate[int, T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return RemoveIndexIf(l, predicate)
}

func (l *List[T]) KeepIf(predicate util.Predicate[T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return KeepIf(l, predicate)
}

func (l *List[T]) KeepIndexIf(predicate util.BiPredicate[int, T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return KeepIndexIf(l, predicate)
}
