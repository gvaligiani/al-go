package dict

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert to set

// ConvertToSet converts a dict of items into a set of other items
func ConvertToSet[K comparable, V any, O comparable](items map[K]V, convert fn.BiConverter[K, V, O]) map[O]struct{} {
	converted := make(map[O]struct{}, len(items))
	for key, value := range items {
		converted[convert(key, value)] = struct{}{}
	}
	return converted
}
