package fn

// //////////////////////////////////////////////////
// supplier

// A Supplier is a function that generates one object
type Supplier[V any] func() V

// A BiSupplier is a function that generates two objets
type BiSupplier[U any, V any] func() (U, V)
