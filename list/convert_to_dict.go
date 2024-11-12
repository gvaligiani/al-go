package list

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert to dict

// ConvertToDict converts a list of items into a dict of other items
func ConvertToDict[T any, K comparable, V any](items []T, convert fn.ToBiConverter[T, K, V]) map[K]V {
	converted := make(map[K]V, len(items))
	for _, item := range items {
		key, value := convert(item)
		converted[key] = value
	}
	return converted
}
