package dict

import "github.com/gvaligiani/al-go/fn"

// //////////////////////////////////////////////////
// keep valid

// KeepValid triggers validation of each item and keeps only the valid ones.
//
// The error is silently discarded, unless 'onError' consumer is provided.
func KeepValid[K comparable, V any](items map[K]V, validate fn.BiValidator[K, V], onError fn.BiErrConsumer[K, V]) map[K]V {
	decodeFn := func(key K, value V) (K, V, error) {
		if err := validate(key, value); err != nil {
			return key, value, err
		}
		return key, value, nil
	}
	return Decode(items, decodeFn, onError)
}
