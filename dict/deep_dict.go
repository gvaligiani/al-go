package dict

import "github.com/gvaligiani/al.go/util"

// alias

type DeepDict[K comparable, V any] map[K]V

// builder

func NewDeep[K comparable, V any]() DeepDict[K, V] {
	return DeepDict[K, V]{}
}

func (d *DeepDict[K, V]) With(key K, value V) *DeepDict[K, V] {
	d.Add(key, value)
	return d
}

// getter

func (d DeepDict[K, V]) Keys() []K {
	l := make([]K, 0, len(d))
	for key := range d {
		l = append(l, key)
	}
	return l
}

func (d DeepDict[K, V]) Values() []V {
	l := make([]V, 0, len(d))
	for _, v := range d {
		l = append(l, v)
	}
	return l
}

// state

func (d DeepDict[K, V]) IsEmpty() bool {
	return len(d) == 0
}

func (d DeepDict[K, V]) AllOf(predicate util.Predicate[V]) bool {
	return AllOf(d, predicate)
}

func (d DeepDict[K, V]) AllKeyOf(predicate util.BiPredicate[K, V]) bool {
	return AllKeyOf(d, predicate)
}

func (d DeepDict[K, V]) AnyOf(predicate util.Predicate[V]) bool {
	return AnyOf(d, predicate)
}

func (d DeepDict[K, V]) AnyKeyOf(predicate util.BiPredicate[K, V]) bool {
	return AnyKeyOf(d, predicate)
}

func (d DeepDict[K, V]) NoneOf(predicate util.Predicate[V]) bool {
	return NoneOf(d, predicate)
}

func (d DeepDict[K, V]) NoKeyOf(predicate util.BiPredicate[K, V]) bool {
	return NoKeyOf(d, predicate)
}

// each

func (d DeepDict[K, V]) Each(consumer util.Consumer[V]) {
	Each(d, consumer)
}

func (d DeepDict[K, V]) EachKey(consumer util.BiConsumer[K, V]) {
	EachKey(d, consumer)
}

// find

func (d DeepDict[K, V]) FindKey(key K) bool {
	return FindKey(d, key)
}

func (d DeepDict[K, V]) FindValueFromKey(key K) (V, bool) {
	return FindValueFromKey(d, key)
}

func (d DeepDict[K, V]) Find(value V) bool {
	return DeepFind(d, value)
}

func (d DeepDict[K, V]) FindKeyFromValue(value V) (K, bool) {
	return DeepFindKeyFromValue(d, value)
}

func (d DeepDict[K, V]) FindIf(predicate util.Predicate[V]) (V, bool) {
	return FindIf(d, predicate)
}

func (d DeepDict[K, V]) FindIfKey(predicate util.BiPredicate[K, V]) (K, V, bool) {
	return FindIfKey(d, predicate)
}

func (d DeepDict[K, V]) FindIfNot(predicate util.Predicate[V]) (V, bool) {
	return FindIfNot(d, predicate)
}

func (d DeepDict[K, V]) FindIfNotKey(predicate util.BiPredicate[K, V]) (K, V, bool) {
	return FindIfNotKey(d, predicate)
}

// copy

func (d DeepDict[K, V]) Copy() DeepDict[K, V] {
	return DeepDict[K, V](Copy(d))
}

func (d DeepDict[K, V]) CopyIf(predicate util.Predicate[V]) DeepDict[K, V] {
	return DeepDict[K, V](CopyIf(d, predicate))
}

func (d DeepDict[K, V]) CopyKeyIf(predicate util.BiPredicate[K, V]) DeepDict[K, V] {
	return DeepDict[K, V](CopyIfKey(d, predicate))
}

func (d DeepDict[K, V]) CopyIfNot(predicate util.Predicate[V]) DeepDict[K, V] {
	return DeepDict[K, V](CopyIfNot(d, predicate))
}

func (d DeepDict[K, V]) CopyKeyIfNot(predicate util.BiPredicate[K, V]) DeepDict[K, V] {
	return DeepDict[K, V](CopyIfNotKey(d, predicate))
}

// modifier

func (d *DeepDict[K, V]) Add(key K, value V) bool {
	_, overriden := (*d)[key]
	(*d)[key] = value
	return overriden
}

func (d *DeepDict[K, V]) Remove(key K) bool {
	if _, ok := (*d)[key]; ok {
		delete(*d, key)
		return true
	}
	return false
}

func (d *DeepDict[K, V]) Clear() bool {
	if d == nil || len(*d) == 0 {
		return false
	}
	*d = DeepDict[K, V]{}
	return true
}

func (d *DeepDict[K, V]) RemoveIf(predicate util.Predicate[V]) bool {
	return RemoveIf(d, predicate)
}

func (d *DeepDict[K, V]) RemoveIfKey(predicate util.BiPredicate[K, V]) bool {
	return RemoveIfKey(d, predicate)
}

func (d *DeepDict[K, V]) KeepIf(predicate util.Predicate[V]) bool {
	return KeepIf(d, predicate)
}

func (d *DeepDict[K, V]) KeepIfKey(predicate util.BiPredicate[K, V]) bool {
	return KeepIfKey(d, predicate)
}
