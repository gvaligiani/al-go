package util

// alias

type Predicate[T any] func(T) bool
type BiPredicate[T any, U any] func(T, U) bool

// predicate <-> not predicate

func Not[T any](predicate Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return !predicate(t)
	}
}

func BiNot[T any, U any](predicate BiPredicate[T, U]) BiPredicate[T, U] {
	return func(t T, u U) bool {
		return !predicate(t, u)
	}
}

// true predicate

func True[T any]() Predicate[T] {
	return func(_ T) bool {
		return true
	}
}

func BiTrue[T any, U any]() BiPredicate[T, U] {
	return func(_ T, _ U) bool {
		return true
	}
}

// false predicate

func False[T any]() Predicate[T] {
	return func(_ T) bool {
		return false
	}
}

func BiFalse[T any, U any]() BiPredicate[T, U] {
	return func(_ T, _ U) bool {
		return false
	}
}

// predicate <-> bi-predicate

func TestWithDefaultFirstArg[T any, U any](predicate BiPredicate[T, U]) Predicate[T] {
	return func(t T) bool {
		var u U
		return predicate(t, u)
	}
}

func TestOnFirstArg[T any, U any](predicate Predicate[T]) BiPredicate[T, U] {
	return func(t T, _ U) bool {
		return predicate(t)
	}
}

func TestWithDefaultSecondArg[T any, U any](predicate BiPredicate[T, U]) Predicate[U] {
	return func(u U) bool {
		var t T
		return predicate(t, u)
	}
}

func TestOnSecondArg[T any, U any](predicate Predicate[U]) BiPredicate[T, U] {
	return func(_ T, u U) bool {
		return predicate(u)
	}
}
