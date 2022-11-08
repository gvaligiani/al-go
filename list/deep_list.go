package list

import "reflect"

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

func (l DeepList[T]) AllOf(predicate Predicate[T]) bool {
	return AllOf(l, predicate)
}

func (l DeepList[T]) AnyOf(predicate Predicate[T]) bool {
	return AnyOf(l, predicate)
}

func (l DeepList[T]) NoneOf(predicate Predicate[T]) bool {
	return NoneOf(l, predicate)
}

// each

func (l DeepList[T]) Each(consumer Consumer[T]) {
	Each(l, consumer)
}

// find

func (l DeepList[T]) Find(value T) (int, bool) {
	return DeepFind(l, value)
}

func (l DeepList[T]) FindIf(predicate Predicate[T]) (int, T, bool) {
	return FindIf(l, predicate)
}

func (l DeepList[T]) FindIfNot(predicate Predicate[T]) (int, T, bool) {
	return FindIfNot(l, predicate)
}

// copy

func (l DeepList[T]) Copy() DeepList[T] {
	return DeepList[T](Copy(l))
}

func (l DeepList[T]) CopyIf(predicate Predicate[T]) DeepList[T] {
	return DeepList[T](CopyIf(l, predicate))
}

func (l DeepList[T]) CopyIfNot(predicate Predicate[T]) DeepList[T] {
	return DeepList[T](CopyIfNot(l, predicate))
}

// modifier

func (l *DeepList[T]) Add(values ...T) bool {
	*l = append(*l, values...)
	return true
}

func (l *DeepList[T]) Remove(value T) bool {
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

func (l *DeepList[T]) Clear() {
	*l = DeepList[T]{}
}

func (l *DeepList[T]) RemoveIf(predicate Predicate[T]) {
	RemoveIf(l, predicate)
}

func (l *DeepList[T]) KeepIf(predicate Predicate[T]) {
	KeepIf(l, predicate)
}
