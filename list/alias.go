package list

import "reflect"

// alias

type List[T any] []T
type Predicate[T any] func(int, T) bool
type Consumer[T any] func(int, T)
type Supplier[T any] func() T
type Transformer[T any, O any] func(int, T) O

// predicate

func Not[T any](predicate Predicate[T]) Predicate[T] {
	return func(index int, item T) bool {
		return !predicate(index, item)
	}
}

func True[T any]() Predicate[T] {
	return func(_ int, _ T) bool {
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

func (l List[T]) Find(value T) bool {
	return Find(l, value)
}

func (l List[T]) FindIf(predicate Predicate[T]) (T, bool) {
	return FindIf(l, predicate)
}

func (l List[T]) FindIfNot(predicate Predicate[T]) (T, bool) {
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

func (l *List[T]) RemoveIf(predicate Predicate[T]) {
	RemoveIf(l, predicate)
}

func (l *List[T]) KeepIf(predicate Predicate[T]) {
	KeepIf(l, predicate)
}
