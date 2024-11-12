package list

import (
	"github.com/gvaligiani/al-go/fn"
	"golang.org/x/exp/constraints"
)

// //////////////////////////////////////////////////
// each

// Each iterates through an unordered list
func Each[V any](l []V, consume fn.Consumer[V]) {
	for _, v := range l {
		consume(v)
	}
}

// //////////////////////////////////////////////////
// each sorted

// EachSorted iterates through a ordered list
func EachSorted[V constraints.Ordered](l []V, consume fn.Consumer[V]) {
	Sort(l)
	for _, v := range l {
		consume(v)
	}
}
