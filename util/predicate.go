package util

// alias

type Predicate[V any] func(V) bool
type BiPredicate[U any, V any] func(U, V) bool

// predicate <-> not predicate

func Not[V any](predicate Predicate[V]) Predicate[V] {
	return func(v V) bool {
		return !predicate(v)
	}
}

func BiNot[U any, V any](predicate BiPredicate[U, V]) BiPredicate[U, V] {
	return func(u U, v V) bool {
		return !predicate(u, v)
	}
}

// true predicate

func True[V any]() Predicate[V] {
	return func(_ V) bool {
		return true
	}
}

func BiTrue[U any, V any]() BiPredicate[U, V] {
	return func(_ U, _ V) bool {
		return true
	}
}

// false predicate

func False[V any]() Predicate[V] {
	return func(_ V) bool {
		return false
	}
}

func BiFalse[U any, V any]() BiPredicate[U, V] {
	return func(_ U, _ V) bool {
		return false
	}
}

// predicate <-> bi-predicate

func TestWithDefaultFirstArg[U any, V any](predicate BiPredicate[U, V]) Predicate[U] {
	return func(u U) bool {
		var v V
		return predicate(u, v)
	}
}

func TestOnFirstArg[U any, V any](predicate Predicate[U]) BiPredicate[U, V] {
	return func(u U, _ V) bool {
		return predicate(u)
	}
}

func TestWithDefaultSecondArg[U any, V any](predicate BiPredicate[U, V]) Predicate[U] {
	return func(u U) bool {
		var v V
		return predicate(u, v)
	}
}

func TestOnSecondArg[U any, V any](predicate Predicate[V]) BiPredicate[U, V] {
	return func(_ U, v V) bool {
		return predicate(v)
	}
}
