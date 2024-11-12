package dict

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// //////////////////////////////////////////////////
// sorted keys

// SortedKeys extracts map keys as a sorted slice
//
//	>>> helpful to iterate through ordered keys
func SortedKeys[K constraints.Ordered, V any](items map[K]V) []K {
	keys := make([]K, 0, len(items))
	for key := range items {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}
