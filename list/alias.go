package list

import "github.com/gvaligiani/al.go/util"

// alias

type Predicate[T any] util.BiPredicate[int, T]
type Consumer[T any] util.BiConsumer[int, T]
type Supplier[T any] util.Supplier[T]
type Transformer[T any, O any] util.BiTransformer[int, T, O]

// predicate

func Not[T any](predicate Predicate[T]) Predicate[T] {
	return func(index int, item T) bool {
		return !predicate(index, item)
	}
}

func True[T any]() Predicate[T] {
	return func(_ int, _ T) bool {
		return true
	}
}
