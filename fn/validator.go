package fn

// //////////////////////////////////////////////////
// validator

// A Validator is a function that returns an error if one provided object is invalid
type Validator[V any] func(V) error

// A BiValidator is a function that returns an error if two provided objects are invalid
type BiValidator[U any, V any] func(U, V) error
