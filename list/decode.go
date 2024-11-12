package list

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// decode

// Decode converts a list of item into a list of other items
//
// Decoding of each item could lead to an error, that does not stop the process.
//
// The error is silently discarded, unless 'onError' consumer is provided.
func Decode[T any, U any](items []T, decode fn.Decoder[T, U], onError fn.ErrConsumer[T]) []U {
	decoded := make([]U, 0, len(items))
	for _, item := range items {
		value, err := decode(item)
		if err == nil {
			decoded = append(decoded, value)
		} else if onError != nil {
			onError(item, err)
		}
	}
	return decoded
}
