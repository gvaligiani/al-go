package fn

// //////////////////////////////////////////////////
// converter

// A Converter is a function that converts one object into one other object
type Converter[V any, O any] func(V) O

// A BiConverter is a function that converts two objects into one other object
type BiConverter[U any, V any, O any] func(U, V) O

// A ToBiConverter is a function that converts one object into two other objects
type ToBiConverter[U any, O any, P any] func(U) (O, P)

// A BiToBiConverter is a function that converts two objects into two other objects
type BiToBiConverter[U any, V any, O any, P any] func(U, V) (O, P)

// //////////////////////////////////////////////////
// helper

// Identity is a specific converter that keeps the same type
func Identity[T any](t T) T {
	return t
}

// BiIdentity is a specific converter that keeps the same two types
func BiIdentity[K any, V any](k K, v V) (K, V) {
	return k, v
}
