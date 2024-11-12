package dict

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// set

type Dict[K comparable, V any] map[K]V

// //////////////////////////////////////////////////
// has

// Has is an helper to test if a key is present in the map
func (d Dict[K, V]) Has(key K) bool {
	_, found := d[key]
	return found
}

// //////////////////////////////////////////////////
// add

// Add is an helper to add a single pair in the map
func (d *Dict[K, V]) Add(key K, value V) {
	(*d)[key] = value
}

// AddPair is an helper to add multiple pairs into the map
func (d *Dict[K, V]) AddPair(items ...pair[K, V]) {
	for _, item := range items {
		(*d)[item.Key] = item.Value
	}
}

// //////////////////////////////////////////////////
// del

// Del is an helper to remove multiple keys from the map
func (d *Dict[K, V]) Del(keys ...K) {
	for _, key := range keys {
		delete((*d), key)
	}
}

// //////////////////////////////////////////////////
// keys

// Keys is an helper to get keys
func (d Dict[K, V]) Keys() []K {
	return Keys(d)
}

// KeySet is an helper to get keys as a set
func (d Dict[K, V]) KeySet() map[K]struct{} {
	return KeySet(d)
}

// //////////////////////////////////////////////////
// state

func (d Dict[K, V]) AllOf(predicate fn.BiPredicate[K, V]) bool {
	return AllOf(d, predicate)
}

func (d Dict[K, V]) AnyOf(predicate fn.BiPredicate[K, V]) bool {
	return AnyOf(d, predicate)
}

func (d Dict[K, V]) NoneOf(predicate fn.BiPredicate[K, V]) bool {
	return NoneOf(d, predicate)
}

// //////////////////////////////////////////////////
// remove if

func (d *Dict[K, V]) RemoveIf(predicate fn.BiPredicate[K, V]) {
	RemoveIf(d, predicate)
}

// //////////////////////////////////////////////////
// keep if

func (d *Dict[K, V]) KeepIf(predicate fn.BiPredicate[K, V]) {
	KeepIf(d, predicate)
}

// //////////////////////////////////////////////////
// count if

func (d Dict[K, V]) CountIf(predicate fn.BiPredicate[K, V]) int {
	return CountIf(d, predicate)
}

// //////////////////////////////////////////////////
// each

func (d Dict[K, V]) Each(consume fn.BiConsumer[K, V]) {
	Each(d, consume)
}

// //////////////////////////////////////////////////
// copy if

func (d Dict[K, V]) CopyIf(predicate fn.BiPredicate[K, V]) Dict[K, V] {
	return Dict[K, V](CopyIf(d, predicate))
}

// //////////////////////////////////////////////////
// find if

func (d Dict[K, V]) FindIf(predicate fn.BiPredicate[K, V]) (K, V, bool) {
	return FindIf(d, predicate)
}

// //////////////////////////////////////////////////
// join

func (d Dict[K, V]) Join(separator string) string {
	return Join(d, separator)
}

func (d Dict[K, V]) ConvertAndJoin(convert fn.BiConverter[K, V, string], separator string) string {
	return ConvertAndJoin(d, convert, separator)
}

// //////////////////////////////////////////////////
// validate

func (d Dict[K, V]) Validate(validate fn.BiValidator[K, V]) error {
	return Validate(d, validate)
}
