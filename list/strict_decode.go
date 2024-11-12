package list

import (
	"github.com/gvaligiani/al-go/fn"
)

// //////////////////////////////////////////////////
// strict decode

// StrictDecode converts a list of item into a list of other items
//
// Decoding of each item could lead to an error, that stops the process
func StrictDecode[T any, U any](items []T, decode fn.Decoder[T, U]) ([]U, error) {
	decoded := make([]U, 0, len(items))
	for _, item := range items {
		value, err := decode(item)
		if err != nil {
			return nil, err
		}
		decoded = append(decoded, value)
	}
	return decoded, nil
}
