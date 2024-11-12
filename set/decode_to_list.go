package set

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// decode to list

// Decode converts a set of items into a list of other items
//
// Decoding of each item could lead to an error, that does not stop the process.
//
// The error is silently discarded, unless 'onError' consumer is provided.
func DecodeToList[T comparable, U any](items map[T]struct{}, decode fn.Decoder[T, U], onError fn.ErrConsumer[T]) []U {
	decoded := make([]U, 0, len(items))
	for item := range items {
		value, err := decode(item)
		if err == nil {
			decoded = append(decoded, value)
		} else if onError != nil {
			onError(item, err)
		}
	}
	return decoded
}
