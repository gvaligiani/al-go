package list

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// keep valid

// KeepValid triggers validation of each item and keeps only the valid ones.
//
// The error is silently discarded, unless 'onError' consumer is provided.
func KeepValid[T any](items []T, validate fn.Validator[T], onError fn.ErrConsumer[T]) []T {
	decodeFn := func(item T) (T, error) {
		if err := validate(item); err != nil {
			return item, err
		}
		return item, nil
	}
	return Decode(items, decodeFn, onError)
}
