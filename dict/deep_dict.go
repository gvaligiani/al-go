package dict

import "github.com/gvaligiani/al.go/util"

// alias

type DeepDict[K comparable, T any] map[K]T

// builder

func NewDeep[K comparable, T any]() DeepDict[K, T] {
	return DeepDict[K, T]{}
}

// getter

func (d DeepDict[K, T]) Keys() []K {
	l := make([]K, 0, len(d))
	for key := range d {
		l = append(l, key)
	}
	return l
}

func (d DeepDict[K, T]) Values() []T {
	l := make([]T, 0, len(d))
	for _, item := range d {
		l = append(l, item)
	}
	return l
}

// state

func (d DeepDict[K, T]) IsEmpty() bool {
	return len(d) == 0
}

func (d DeepDict[K, T]) AllOf(predicate util.Predicate[T]) bool {
	return AllOf(d, predicate)
}

func (d DeepDict[K, T]) AllKeyOf(predicate util.BiPredicate[K, T]) bool {
	return AllKeyOf(d, predicate)
}

func (d DeepDict[K, T]) AnyOf(predicate util.Predicate[T]) bool {
	return AnyOf(d, predicate)
}

func (d DeepDict[K, T]) AnyKeyOf(predicate util.BiPredicate[K, T]) bool {
	return AnyKeyOf(d, predicate)
}

func (d DeepDict[K, T]) NoneOf(predicate util.Predicate[T]) bool {
	return NoneOf(d, predicate)
}

func (d DeepDict[K, T]) NoKeyOf(predicate util.BiPredicate[K, T]) bool {
	return NoKeyOf(d, predicate)
}

// each

func (d DeepDict[K, T]) Each(consumer util.Consumer[T]) {
	Each(d, consumer)
}

func (d DeepDict[K, T]) EachKey(consumer util.BiConsumer[K, T]) {
	EachKey(d, consumer)
}

// find

func (d DeepDict[K, T]) FindKey(key K) bool {
	return FindKey(d, key)
}

func (d DeepDict[K, T]) FindValueOf(key K) (T, bool) {
	return FindValueOf(d, key)
}

func (d DeepDict[K, T]) Find(value T) bool {
	return DeepFind(d, value)
}

func (d DeepDict[K, T]) FindKeyOf(value T) (K, bool) {
	return DeepFindKeyOf(d, value)
}

func (d DeepDict[K, T]) FindIf(predicate util.Predicate[T]) (T, bool) {
	return FindIf(d, predicate)
}

func (d DeepDict[K, T]) FindKeyIf(predicate util.BiPredicate[K, T]) (K, T, bool) {
	return FindKeyIf(d, predicate)
}

func (d DeepDict[K, T]) FindIfNot(predicate util.Predicate[T]) (T, bool) {
	return FindIfNot(d, predicate)
}

func (d DeepDict[K, T]) FindKeyIfNot(predicate util.BiPredicate[K, T]) (K, T, bool) {
	return FindKeyIfNot(d, predicate)
}

// copy

func (d DeepDict[K, T]) Copy() DeepDict[K, T] {
	return DeepDict[K, T](Copy(d))
}

func (d DeepDict[K, T]) CopyIf(predicate util.Predicate[T]) DeepDict[K, T] {
	return DeepDict[K, T](CopyIf(d, predicate))
}

func (d DeepDict[K, T]) CopyKeyIf(predicate util.BiPredicate[K, T]) DeepDict[K, T] {
	return DeepDict[K, T](CopyKeyIf(d, predicate))
}

func (d DeepDict[K, T]) CopyIfNot(predicate util.Predicate[T]) DeepDict[K, T] {
	return DeepDict[K, T](CopyIfNot(d, predicate))
}

func (d DeepDict[K, T]) CopyKeyIfNot(predicate util.BiPredicate[K, T]) DeepDict[K, T] {
	return DeepDict[K, T](CopyKeyIfNot(d, predicate))
}

// modifier

func (d *DeepDict[K, T]) Add(key K, item T) bool {
	_, overriden := (*d)[key]
	(*d)[key] = item
	return overriden
}

func (d *DeepDict[K, T]) Remove(key K) bool {
	if _, ok := (*d)[key]; ok {
		delete(*d, key)
		return true
	}
	return false
}

func (d *DeepDict[K, T]) Clear() bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	*d = DeepDict[K, T]{}
	return true
}

func (d *DeepDict[K, T]) RemoveIf(predicate util.Predicate[T]) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	return RemoveIf(d, predicate)
}

func (d *DeepDict[K, T]) RemoveKeyIf(predicate util.BiPredicate[K, T]) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	return RemoveKeyIf(d, predicate)
}

func (d *DeepDict[K, T]) KeepIf(predicate util.Predicate[T]) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	return KeepIf(d, predicate)
}

func (d *DeepDict[K, T]) KeepKeyIf(predicate util.BiPredicate[K, T]) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	return KeepKeyIf(d, predicate)
}
