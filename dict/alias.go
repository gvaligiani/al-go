package dict

import "github.com/gvaligiani/al.go/util"

// alias

type Predicate[K comparable, T any] util.BiPredicate[K, T]
type Consumer[K comparable, T any] util.BiConsumer[K, T]
type Supplier[K comparable, T any] util.BiSupplier[K, T]
type Transformer[K comparable, T any, O any] util.BiTransformer[K, T, O]

// helper

func Not[K comparable, T any](predicate Predicate[K, T]) Predicate[K, T] {
	return func(k K, t T) bool {
		return !predicate(k, t)
	}
}

func True[K comparable, T any]() Predicate[K, T] {
	return func(_ K, _ T) bool {
		return true
	}
}
