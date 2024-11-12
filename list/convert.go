package list

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert

// Convert converts a list of items into a list of other items
func Convert[T any, U any](items []T, convert fn.Converter[T, U]) []U {
	converted := make([]U, 0, len(items))
	for _, item := range items {
		converted = append(converted, convert(item))
	}
	return converted
}
