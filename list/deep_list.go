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

func (l *DeepList[T]) Add(value T) bool {
	*l = append(*l, value)
	return true
}

func (l *DeepList[T]) Remove(value T) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return RemoveIf(l, func(_ int, t T) bool { return util.DeepEqual(t, value) })
}

func (l *DeepList[T]) Clear() bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	*l = DeepList[T]{}
	return true
}

func (l *DeepList[T]) RemoveIf(predicate Predicate[T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return RemoveIf(l, predicate)
}

func (l *DeepList[T]) KeepIf(predicate Predicate[T]) bool {
	if l == nil || len(*l) == 0 {
		return false
	}
	return KeepIf(l, predicate)
}
