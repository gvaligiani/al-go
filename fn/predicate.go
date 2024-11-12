package fn

// //////////////////////////////////////////////////
// predicate

// A Predicate is a function that checks something on one object
type Predicate[V any] func(V) bool

// A BiPredicate is a function that checks something on two objects
type BiPredicate[U any, V any] func(U, V) bool

// //////////////////////////////////////////////////
// reverse predicate

// Not is an helper to revert a [Predicate]
func Not[V any](predicate Predicate[V]) Predicate[V] {
	return func(v V) bool {
		return !predicate(v)
	}
}

// BiNot is an helper to revert a [BiPredicate]
func BiNot[U any, V any](predicate BiPredicate[U, V]) BiPredicate[U, V] {
	return func(u U, v V) bool {
		return !predicate(u, v)
	}
}

// //////////////////////////////////////////////////
// true predicate

// True is a [Predicate] which is always true
func True[V any]() Predicate[V] {
	return func(_ V) bool {
		return true
	}
}

// BiTrue is a [BiPredicate] which is always true
func BiTrue[U any, V any]() BiPredicate[U, V] {
	return func(_ U, _ V) bool {
		return true
	}
}

// //////////////////////////////////////////////////
// false predicate

// False is a [Predicate] which is always false
func False[V any]() Predicate[V] {
	return func(_ V) bool {
		return false
	}
}

// BiFalse is a [BiPredicate] which is always false
func BiFalse[U any, V any]() BiPredicate[U, V] {
	return func(_ U, _ V) bool {
		return false
	}
}
