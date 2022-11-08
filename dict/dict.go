package dict

import "github.com/gvaligiani/al.go/util"

// alias

type Dict[K comparable, V comparable] map[K]V

// builder

func New[K comparable, V comparable]() Dict[K, V] {
	// TODO add capacity
	return Dict[K, V]{}
}

func (d Dict[K, V]) With(key K, value V) Dict[K, V] {
	d.Add(key, value)
	return d
}

// getter

func (d Dict[K, V]) Keys() []K {
	l := make([]K, 0, len(d))
	for key := range d {
		l = append(l, key)
	}
	return l
}

func (d Dict[K, V]) Values() []V {
	l := make([]V, 0, len(d))
	for _, v := range d {
		l = append(l, v)
	}
	return l
}

// state

func (d Dict[K, V]) IsEmpty() bool {
	return len(d) == 0
}

func (d Dict[K, V]) AllOf(predicate util.Predicate[V]) bool {
	return AllOf(d, predicate)
}

func (d Dict[K, V]) AllKeyOf(predicate util.BiPredicate[K, V]) bool {
	return AllKeyOf(d, predicate)
}

func (d Dict[K, V]) AnyOf(predicate util.Predicate[V]) bool {
	return AnyOf(d, predicate)
}

func (d Dict[K, V]) AnyKeyOf(predicate util.BiPredicate[K, V]) bool {
	return AnyKeyOf(d, predicate)
}

func (d Dict[K, V]) NoneOf(predicate util.Predicate[V]) bool {
	return NoneOf(d, predicate)
}

func (d Dict[K, V]) NoKeyOf(predicate util.BiPredicate[K, V]) bool {
	return NoKeyOf(d, predicate)
}

// each

func (d Dict[K, V]) Each(consumer util.Consumer[V]) {
	Each(d, consumer)
}

func (d Dict[K, V]) EachKey(consumer util.BiConsumer[K, V]) {
	EachKey(d, consumer)
}

// find

func (d Dict[K, V]) FindKey(key K) bool {
	return FindKey(d, key)
}

func (d Dict[K, V]) FindValueFromKey(key K) (V, bool) {
	return FindValueFromKey(d, key)
}

func (d Dict[K, V]) Find(value V) bool {
	return Find(d, value)
}

func (d Dict[K, V]) FindKeyFromValue(value V) (K, bool) {
	return FindKeyFromValue(d, value)
}

func (d Dict[K, V]) FindIf(predicate util.Predicate[V]) (V, bool) {
	return FindIf(d, predicate)
}

func (d Dict[K, V]) FindIfKey(predicate util.BiPredicate[K, V]) (K, V, bool) {
	return FindIfKey(d, predicate)
}

func (d Dict[K, V]) FindIfNot(predicate util.Predicate[V]) (V, bool) {
	return FindIfNot(d, predicate)
}

func (d Dict[K, V]) FindIfNotKey(predicate util.BiPredicate[K, V]) (K, V, bool) {
	return FindIfNotKey(d, predicate)
}

// copy

func (d Dict[K, V]) Copy() Dict[K, V] {
	return Dict[K, V](Copy(d))
}

func (d Dict[K, V]) CopyIf(predicate util.Predicate[V]) Dict[K, V] {
	return Dict[K, V](CopyIf(d, predicate))
}

func (d Dict[K, V]) CopyIfKey(predicate util.BiPredicate[K, V]) Dict[K, V] {
	return Dict[K, V](CopyIfKey(d, predicate))
}

func (d Dict[K, V]) CopyIfNot(predicate util.Predicate[V]) Dict[K, V] {
	return Dict[K, V](CopyIfNot(d, predicate))
}

func (d Dict[K, V]) CopyIfNotKey(predicate util.BiPredicate[K, V]) Dict[K, V] {
	return Dict[K, V](CopyIfNotKey(d, predicate))
}

// modifier

func (d *Dict[K, V]) Add(key K, value V) bool {
	_, overriden := (*d)[key]
	(*d)[key] = value
	return overriden
}

func (d *Dict[K, V]) Remove(key K) bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	if _, ok := (*d)[key]; ok {
		delete(*d, key)
		return true
	}
	return false
}

func (d *Dict[K, V]) Clear() bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	*d = Dict[K, V]{}
	return true
}

func (d *Dict[K, V]) RemoveIf(predicate util.Predicate[V]) bool {
	return RemoveIf(d, predicate)
}

func (d *Dict[K, V]) RemoveIfKey(predicate util.BiPredicate[K, V]) bool {
	return RemoveIfKey(d, predicate)
}

func (d *Dict[K, V]) KeepIf(predicate util.Predicate[V]) bool {
	return KeepIf(d, predicate)
}

func (d *Dict[K, V]) KeepIfKey(predicate util.BiPredicate[K, V]) bool {
	return KeepIfKey(d, predicate)
}
