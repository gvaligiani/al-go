package set

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// decode

// Decode converts a set of items into a set of other items
//
// Decoding of each item could lead to an error, that does not stop the process.
//
// The error is silently discarded, unless 'onError' consumer is provided.
func Decode[T comparable, U comparable](items map[T]struct{}, decode fn.Decoder[T, U], onError fn.ErrConsumer[T]) map[U]struct{} {
	decoded := make(map[U]struct{}, len(items))
	for item := range items {
		value, err := decode(item)
		if err == nil {
			decoded[value] = struct{}{}
		} else if onError != nil {
			onError(item, err)
		}
	}
	return decoded
}
