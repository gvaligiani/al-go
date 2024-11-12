package list

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// //////////////////////////////////////////////////
// sort

// Sort is an helper to sort a simple list
func Sort[T constraints.Ordered](items []T) {
	sort.Slice(items, func(i, j int) bool { return items[i] < items[j] })
}
