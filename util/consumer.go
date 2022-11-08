package util

// alias

type Consumer[T any] func(T)
type BiConsumer[T any, U any] func(T, U)

// consumer <-> bi-consumer

func ConsumeWithDefaultFirstArg[T any, U any](consumer BiConsumer[T, U]) Consumer[T] {
	return func(t T) {
		var u U
		consumer(t, u)
	}
}

func ConsumeOnFirstArg[T any, U any](consumer Consumer[T]) BiConsumer[T, U] {
	return func(t T, _ U) {
		consumer(t)
	}
}

func ConsumeWithDefaultSecondArg[T any, U any](consumer BiConsumer[T, U]) Consumer[U] {
	return func(u U) {
		var t T
		consumer(t, u)
	}
}

func ConsumeOnSecondArg[T any, U any](consumer Consumer[U]) BiConsumer[T, U] {
	return func(_ T, u U) {
		consumer(u)
	}
}
