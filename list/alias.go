package list

import "reflect"

// alias

type List[T any] []T
type Predicate[T any] func(T) bool
type Consumer[T any] func(T)
type Supplier[T any] func() T
type Transformer[T any, O any] func(T) O

// predicate

func Not[T any](predicate Predicate[T]) Predicate[T] {
	return func(item T) bool {
		return !predicate(item)
	}
}

func True[T any]() Predicate[T] {
	return func(_ T) bool {
		return true
	}
}

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

func (l List[T]) AllOf(predicate func(T) bool) bool {
	return AllOf(l, predicate)
}

func (l List[T]) AnyOf(predicate func(T) bool) bool {
	return AnyOf(l, predicate)
}

func (l List[T]) NoneOf(predicate func(T) bool) bool {
	return NoneOf(l, predicate)
}

// each

func (l List[T]) Each(consumer func(T)) {
	Each(l, consumer)
}

// TODO needed ?
func (l List[T]) EachIndex(consumer func(int, T)) {
	EachIndex(l, consumer)
}

// find

func (l List[T]) Find(value T) bool {
	return Find(l, value)
}

func (l List[T]) FindIf(predicate func(T) bool) (T, bool) {
	return FindIf(l, predicate)
}

func (l List[T]) FindIfNot(predicate func(T) bool) (T, bool) {
	return FindIfNot(l, predicate)
}

// copy

func (l List[T]) Copy() List[T] {
	return List[T](Copy(l))
}

func (l List[T]) CopyIf(predicate func(T) bool) List[T] {
	return List[T](CopyIf(l, predicate))
}

func (l List[T]) CopyIfNot(predicate func(T) bool) List[T] {
	return List[T](CopyIfNot(l, predicate))
}

// modifier

func (l *List[T]) Add(values ...T) bool {
	*l = append(*l, values...)
	return true
}

func (l *List[T]) Remove(value T) bool {
	newL := make([]T, 0, len(*l))
	atLeastOneRemoved := false
	for _, item := range *l {
		if reflect.DeepEqual(item, value) {
			atLeastOneRemoved = true
		} else {
			newL = append(newL, item)
		}
	}
	*l = newL
	return atLeastOneRemoved
}

func (l *List[T]) Clear() {
	*l = []T{}
}

func (l *List[T]) RemoveIf(predicate func(T) bool) {
	RemoveIf(l, predicate)
}

func (l *List[T]) KeepIf(predicate func(T) bool) {
	KeepIf(l, predicate)
}
