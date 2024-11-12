package fn

// //////////////////////////////////////////////////
// consumer

// A Consumer is a function that consumes an object without producing any output
type Consumer[V any] func(V)

// A BiConsumer is a function that consumes two objects without producing any output
type BiConsumer[U any, V any] func(U, V)

// //////////////////////////////////////////////////
// error consumer

// A Consumer is a function that consumes an object with its respective error without producing any output
type ErrConsumer[V any] func(V, error)

// A BiConsumer is a function that consumes two objects with their respective error without producing any output
type BiErrConsumer[U any, V any] func(U, V, error)
