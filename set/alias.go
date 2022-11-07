package set

import "github.com/gvaligiani/al.go/util"

// alias

type Predicate[T comparable] util.Predicate[T]
type Consumer[T comparable] util.Consumer[T]
type Supplier[T comparable] util.Supplier[T]
type Transformer[T comparable, O any] util.Transformer[T, O]

// predicate

func Not[T comparable](predicate Predicate[T]) Predicate[T] {
	return func(item T) bool {
		return !predicate(item)
	}
}

func True[T comparable]() Predicate[T] {
	return func(_ T) bool {
		return true
	}
}
