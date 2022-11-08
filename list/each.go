package list

import "github.com/gvaligiani/al.go/util"

func Each[T any, L ~[]T](items L, consumer util.Consumer[T]) {
	EachIndex(items, util.ConsumeOnSecondArg[int, T](consumer))
}

func EachIndex[T any, L ~[]T](items L, consumer util.BiConsumer[int, T]) {
	for index, item := range items {
		consumer(index, item)
	}
}
