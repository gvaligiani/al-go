package dict

import "github.com/gvaligiani/al.go/util"

func Each[K comparable, T any, M ~map[K]T](items M, consumer util.Consumer[T]) {
	EachKey(items, util.ConsumeOnSecondArg[K, T](consumer))
}

func EachKey[K comparable, T any, M ~map[K]T](items M, consumer util.BiConsumer[K, T]) {
	for key, item := range items {
		consumer(key, item)
	}
}
