package fn

// //////////////////////////////////////////////////
// reducer

// A Reducer is a function that aggregates one object into one other object
type Reducer[V any, O any] func(V, O) O

// A BiReducer is a function that aggregates two objects into one other object
type BiReducer[U any, V any, O any] func(U, V, O) O
