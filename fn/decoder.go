package fn

// //////////////////////////////////////////////////
// decoder

// A Decoder is a function that decodes one object into one other object ( with potential decoding error )
type Decoder[V any, O any] func(V) (O, error)

// A BiDecoder is a function that decodes two objects into one other object ( with potential decoding error )
type BiDecoder[U any, V any, O any] func(U, V) (O, error)

// A Decoder is a function that decodes two objects into two other objects ( with potential decoding error )
type Bi2BiDecoder[U any, V any, O any, P any] func(U, V) (O, P, error)
