package set

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert to dict

// ConvertToDict converts a set of items into a dict of other items
func ConvertToDict[T comparable, K comparable, V any](items map[T]struct{}, convert fn.ToBiConverter[T, K, V]) map[K]V {
	converted := make(map[K]V, len(items))
	for item := range items {
		key, value := convert(item)
		converted[key] = value
	}
	return converted
}
