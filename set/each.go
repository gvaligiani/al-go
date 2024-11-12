package set

import (
	"github.com/gvaligiani/al-go/fn"
	"golang.org/x/exp/constraints"
)

// //////////////////////////////////////////////////
// each

// Each iterates through all set values
func Each[T comparable](items map[T]struct{}, consume fn.Consumer[T]) {
	for item := range items {
		consume(item)
	}
}

// //////////////////////////////////////////////////
// each sorted

// EachSorted iterates through all set values by order
func EachSorted[T constraints.Ordered](items map[T]struct{}, consume fn.Consumer[T]) {
	for _, item := range ToSorted(items) {
		consume(item)
	}
}
