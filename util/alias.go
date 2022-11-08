package util

//
// consumer
//

type Supplier[V any] func() V
type BiSupplier[U any, V any] func() (U, V)

//
// transformer
//

type Transformer[V any, O any] func(V) O
type BiTransformer[U any, V any, O any] func(U, V) O
