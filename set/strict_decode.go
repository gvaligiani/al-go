package set

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// strict decode

// StrictDecode converts a set of items into a set of other items
//
// Decoding of each item could lead to an error, that stops the process.
func StrictDecode[T comparable, U comparable](items map[T]struct{}, decode fn.Decoder[T, U]) (map[U]struct{}, error) {
	decoded := make(map[U]struct{}, len(items))
	for item := range items {
		value, err := decode(item)
		if err != nil {
			return nil, err
		}
		decoded[value] = struct{}{}
	}
	return decoded, nil
}
