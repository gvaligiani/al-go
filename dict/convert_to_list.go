package dict

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert to list

// ConvertToList converts a dict of items into a list of other items
func ConvertToList[K comparable, V any, U any](items map[K]V, convert fn.BiConverter[K, V, U]) []U {
	converted := make([]U, 0, len(items))
	for key, value := range items {
		converted = append(converted, convert(key, value))
	}
	return converted
}
