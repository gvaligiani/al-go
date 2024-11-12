package list

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert not nil

// ConvertNotNil converts a list of items into a list of other items, but skipping nil values
func ConvertNotNil[T any, U any](items []T, convert fn.Converter[T, *U]) []*U {
	converted := make([]*U, 0, len(items))
	for _, item := range items {
		u := convert(item)
		if u != nil {
			converted = append(converted, u)
		}
	}
	return converted
}
