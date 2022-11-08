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

func (l List[T]) AllOf(predicate Predicate[T]) bool {
	return AllOf(l, predicate)
}

func (l List[T]) AnyOf(predicate Predicate[T]) bool {
	return AnyOf(l, predicate)
}

func (l List[T]) NoneOf(predicate Predicate[T]) bool {
	return NoneOf(l, predicate)
}

// each

func (l List[T]) Each(consumer Consumer[T]) {
	Each(l, consumer)
}

// find

func (l List[T]) Find(value T) (int, bool) {
	return Find(l, value)
}

func (l List[T]) FindIf(predicate Predicate[T]) (int, T, bool) {
	return FindIf(l, predicate)
}

func (l List[T]) FindIfNot(predicate Predicate[T]) (int, T, bool) {
	return FindIfNot(l, predicate)
}

// copy

func (l List[T]) Copy() List[T] {
	return List[T](Copy(l))
}

func (l List[T]) CopyIf(predicate Predicate[T]) List[T] {
	return List[T](CopyIf(l, predicate))
}

func (l List[T]) CopyIfNot(predicate Predicate[T]) List[T] {
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
	return RemoveIf(l, func(_ int, t T) bool { return util.Equal(t, value) })
}

func (l *List[T]) Clear() bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	*l = List[T]{}
	return true
}

func (l *List[T]) RemoveIf(predicate Predicate[T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return RemoveIf(l, predicate)
}

func (l *List[T]) KeepIf(predicate Predicate[T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return KeepIf(l, predicate)
}
