package set

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert

// Convert converts a set of items into a set of other items
func Convert[T comparable, U comparable](items map[T]struct{}, convert fn.Converter[T, U]) map[U]struct{} {
	converted := make(map[U]struct{}, len(items))
	for item := range items {
		converted[convert(item)] = struct{}{}
	}
	return converted
}
