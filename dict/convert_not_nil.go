package dict

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// convert not nil

// ConvertNotNil converts a list of items into a list of other items, but skipping nil values
func ConvertNotNil[K comparable, V any, L comparable, W any](items map[K]V, convert fn.BiToBiConverter[K, V, L, *W]) map[L]*W {
	converted := make(map[L]*W, len(items))
	for key, value := range items {
		newKey, newValue := convert(key, value)
		if newValue != nil {
			converted[newKey] = newValue
		}
	}
	return converted
}
