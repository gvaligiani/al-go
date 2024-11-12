package dict

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// strict decode

// StrictDecode converts a dict of items into a dict of other items
//
// Decoding of each item could lead to an error, that stops the process.
func StrictDecode[K comparable, V any, L comparable, W any](items map[K]V, decode fn.Bi2BiDecoder[K, V, L, W]) (map[L]W, error) {
	decoded := make(map[L]W, len(items))
	for key, value := range items {
		newKey, newValue, err := decode(key, value)
		if err != nil {
			return nil, err
		}
		decoded[newKey] = newValue
	}
	return decoded, nil
}
