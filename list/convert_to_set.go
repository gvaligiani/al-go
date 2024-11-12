package list

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert to set

// Convert converts a list of items into a set of other items
func ConvertToSet[T any, U comparable](items []T, convert fn.Converter[T, U]) map[U]struct{} {
	converted := make(map[U]struct{}, len(items))
	for _, item := range items {
		converted[convert(item)] = struct{}{}
	}
	return converted
}
