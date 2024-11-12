package set

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// keep valid

// KeepValid triggers validation of each item and keeps only the valid ones.
//
// The error is silently discarded, unless 'onError' consumer is provided.
func KeepValid[T comparable](items map[T]struct{}, validate fn.Validator[T], onError fn.ErrConsumer[T]) map[T]struct{} {
	decodeFn := func(item T) (T, error) {
		if err := validate(item); err != nil {
			return item, err
		}
		return item, nil
	}
	return Decode(items, decodeFn, onError)
}
