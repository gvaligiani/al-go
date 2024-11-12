package set

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// //////////////////////////////////////////////////
// to sorted

// ToSorted converts the set into an ordered list
func ToSorted[T constraints.Ordered](items map[T]struct{}) []T {
	values := make([]T, 0, len(items))
	for item := range items {
		values = append(values, item)
	}
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return values
}
