package set

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// decode to list

// StrictDecodeToList converts a set of items into a list of other items
//
// Decoding of each item could lead to an error, that stops the process.
func StrictDecodeToList[T comparable, U any](items map[T]struct{}, decode fn.Decoder[T, U]) ([]U, error) {
	decoded := make([]U, 0, len(items))
	for item := range items {
		value, err := decode(item)
		if err != nil {
			return nil, err
		}
		decoded = append(decoded, value)
	}
	return decoded, nil
}
