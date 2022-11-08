package util

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
