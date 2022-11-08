package util

// alias

type Consumer[V any] func(V)
type BiConsumer[U any, V any] func(U, V)

// consumer <-> bi-consumer

func ConsumeWithDefaultFirstArg[U any, V any](consumer BiConsumer[U, V]) Consumer[U] {
	return func(u U) {
		var v V
		consumer(u, v)
	}
}

func ConsumeOnFirstArg[U any, V any](consumer Consumer[U]) BiConsumer[U, V] {
	return func(u U, _ V) {
		consumer(u)
	}
}

func ConsumeWithDefaultSecondArg[U any, V any](consumer BiConsumer[U, V]) Consumer[V] {
	return func(v V) {
		var u U
		consumer(u, v)
	}
}

func ConsumeOnSecondArg[U any, V any](consumer Consumer[V]) BiConsumer[U, V] {
	return func(_ U, v V) {
		consumer(v)
	}
}
