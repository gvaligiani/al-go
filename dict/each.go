package dict

import (
	"github.com/gvaligiani/al-go/fn"
	"golang.org/x/exp/constraints"
)

// //////////////////////////////////////////////////
// each

// Each iterates through a map by unordered keys
func Each[K comparable, V any](items map[K]V, consume fn.BiConsumer[K, V]) {
	for key, value := range items {
		consume(key, value)
	}
}

// //////////////////////////////////////////////////
// each sorted

// EachSorted iterates through a map by ordered keys
func EachSorted[K constraints.Ordered, V any](items map[K]V, consume fn.BiConsumer[K, V]) {
	for _, key := range SortedKeys(items) {
		consume(key, items[key])
	}
}
