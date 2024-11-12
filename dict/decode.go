package dict

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// decode

// Decode converts a dict of items into a dict of other items
//
// Decoding of each item could lead to an error, that does not stop the process.
//
// The error is silently discarded, unless 'onError' consumer is provided.
func Decode[K comparable, V any, L comparable, W any](items map[K]V, decode fn.Bi2BiDecoder[K, V, L, W], onError fn.BiErrConsumer[K, V]) map[L]W {
	decoded := make(map[L]W, len(items))
	for key, value := range items {
		newKey, newValue, err := decode(key, value)
		if err == nil {
			decoded[newKey] = newValue
		} else if onError != nil {
			onError(key, value, err)
		}
	}
	return decoded
}
