package dict

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert

// Convert converts a dict of items into a dict of other items
func Convert[K comparable, V any, L comparable, W any](items map[K]V, convert fn.BiToBiConverter[K, V, L, W]) map[L]W {
	converted := make(map[L]W, len(items))
	for key, value := range items {
		newKey, newValue := convert(key, value)
		converted[newKey] = newValue
	}
	return converted
}
