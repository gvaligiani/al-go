package util

//
// predicate
//

type Predicate[T any] func(T) bool
type BiPredicate[T any, U any] func(T, U) bool

//
// consumer
//

type Consumer[T any] func(T)
type BiConsumer[T any, U any] func(T, U)

//
// consumer
//

type Supplier[T any] func() T
type BiSupplier[T any, U any] func() (T, U)

//
// transformer
//

type Transformer[T any, O any] func(T) O
type BiTransformer[T any, U any, O any] func(T, U) O
