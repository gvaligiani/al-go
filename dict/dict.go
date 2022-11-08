package dict

import "github.com/gvaligiani/al.go/util"

// alias

type Dict[K comparable, T comparable] map[K]T

// builder

func New[K comparable, T comparable]() Dict[K, T] {
	return Dict[K, T]{}
}

// getter

func (d Dict[K, T]) Keys() []K {
	l := make([]K, 0, len(d))
	for key := range d {
		l = append(l, key)
	}
	return l
}

func (d Dict[K, T]) Values() []T {
	l := make([]T, 0, len(d))
	for _, item := range d {
		l = append(l, item)
	}
	return l
}

// state

func (d Dict[K, T]) IsEmpty() bool {
	return len(d) == 0
}

func (d Dict[K, T]) AllOf(predicate util.Predicate[T]) bool {
	return AllOf(d, predicate)
}

func (d Dict[K, T]) AllKeyOf(predicate util.BiPredicate[K, T]) bool {
	return AllKeyOf(d, predicate)
}

func (d Dict[K, T]) AnyOf(predicate util.Predicate[T]) bool {
	return AnyOf(d, predicate)
}

func (d Dict[K, T]) AnyKeyOf(predicate util.BiPredicate[K, T]) bool {
	return AnyKeyOf(d, predicate)
}

func (d Dict[K, T]) NoneOf(predicate util.Predicate[T]) bool {
	return NoneOf(d, predicate)
}

func (d Dict[K, T]) NoKeyOf(predicate util.BiPredicate[K, T]) bool {
	return NoKeyOf(d, predicate)
}

// each

func (d Dict[K, T]) Each(consumer util.Consumer[T]) {
	Each(d, consumer)
}

func (d Dict[K, T]) EachKey(consumer util.BiConsumer[K, T]) {
	EachKey(d, consumer)
}

// find

func (d Dict[K, T]) FindKey(key K) bool {
	return FindKey(d, key)
}

func (d Dict[K, T]) FindValue(value T) bool {
	return FindValue(d, value)
}

func (d Dict[K, T]) FindIf(predicate util.Predicate[T]) (T, bool) {
	return FindIf(d, predicate)
}

func (d Dict[K, T]) FindKeyIf(predicate util.BiPredicate[K, T]) (K, T, bool) {
	return FindKeyIf(d, predicate)
}

func (d Dict[K, T]) FindIfNot(predicate util.Predicate[T]) (T, bool) {
	return FindIfNot(d, predicate)
}

func (d Dict[K, T]) FindKeyIfNot(predicate util.BiPredicate[K, T]) (K, T, bool) {
	return FindKeyIfNot(d, predicate)
}

// copy

func (d Dict[K, T]) Copy() Dict[K, T] {
	return Dict[K, T](Copy(d))
}

func (d Dict[K, T]) CopyIf(predicate util.Predicate[T]) Dict[K, T] {
	return Dict[K, T](CopyIf(d, predicate))
}

func (d Dict[K, T]) CopyKeyIf(predicate util.BiPredicate[K, T]) Dict[K, T] {
	return Dict[K, T](CopyKeyIf(d, predicate))
}

func (d Dict[K, T]) CopyIfNot(predicate util.Predicate[T]) Dict[K, T] {
	return Dict[K, T](CopyIfNot(d, predicate))
}

func (d Dict[K, T]) CopyKeyIfNot(predicate util.BiPredicate[K, T]) Dict[K, T] {
	return Dict[K, T](CopyKeyIfNot(d, predicate))
}

// modifier

func (d *Dict[K, T]) Add(key K, item T) bool {
	_, overriden := (*d)[key]
	(*d)[key] = item
	return overriden
}

func (d *Dict[K, T]) Remove(key K) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	if _, ok := (*d)[key]; ok {
		delete(*d, key)
		return true
	}
	return false
}

func (d *Dict[K, T]) Clear() bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	*d = Dict[K, T]{}
	return true
}

func (d *Dict[K, T]) RemoveIf(predicate util.Predicate[T]) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	return RemoveIf(d, predicate)
}

func (d *Dict[K, T]) RemoveKeyIf(predicate util.BiPredicate[K, T]) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	return RemoveKeyIf(d, predicate)
}

func (d *Dict[K, T]) KeepIf(predicate util.Predicate[T]) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	return KeepIf(d, predicate)
}

func (d *Dict[K, T]) KeepKeyIf(predicate util.BiPredicate[K, T]) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	return KeepKeyIf(d, predicate)
}
