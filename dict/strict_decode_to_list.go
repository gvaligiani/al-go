package dict

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// strict decode to list

// StrictDecodeToList converts a dict of items into a list of other items
//
// Decoding of each item could lead to an error, that stops the process.
func StrictDecodeToList[K comparable, V any, U any](items map[K]V, decode fn.BiDecoder[K, V, U]) ([]U, error) {
	decoded := make([]U, 0, len(items))
	for key, value := range items {
		newValue, err := decode(key, value)
		if err != nil {
			return nil, err
		}
		decoded = append(decoded, newValue)
	}
	return decoded, nil
}
