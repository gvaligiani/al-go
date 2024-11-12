package dict

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// decode to list

// DecodeToList converts a dict of items into a list of other items
//
// Decoding of each item could lead to an error, that does not stop the process.
//
// The error is silently discarded, unless 'onError' consumer is provided.
func DecodeToList[K comparable, V any, U any](items map[K]V, decode fn.BiDecoder[K, V, U], onError fn.BiErrConsumer[K, V]) []U {
	decoded := make([]U, 0, len(items))
	for key, value := range items {
		newValue, err := decode(key, value)
		if err == nil {
			decoded = append(decoded, newValue)
		} else if onError != nil {
			onError(key, value, err)
		}
	}
	return decoded
}
