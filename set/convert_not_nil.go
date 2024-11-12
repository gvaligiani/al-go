package set

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert not nil

// ConvertNotNil converts a list of items into a list of other items, but skipping nil values
func ConvertNotNil[T comparable, U comparable](items map[T]struct{}, convert fn.Converter[T, *U]) map[*U]struct{} {
	converted := make(map[*U]struct{}, len(items))
	for item := range items {
		u := convert(item)
		if u != nil {
			converted[u] = struct{}{}
		}
	}
	return converted
}
