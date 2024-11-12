package set

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert to list

// ConvertToList converts a set of items into a list of other items
func ConvertToList[T comparable, U any](items map[T]struct{}, convert fn.Converter[T, U]) []U {
	converted := make([]U, 0, len(items))
	for item := range items {
		converted = append(converted, convert(item))
	}
	return converted
}
