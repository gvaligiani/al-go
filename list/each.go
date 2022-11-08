package list

import "github.com/gvaligiani/al.go/util"

func Each[V any, L ~[]V](l L, consumer util.Consumer[V]) {
	EachIndex(l, util.ConsumeOnSecondArg[int](consumer))
}

func EachIndex[V any, L ~[]V](l L, consumer util.BiConsumer[int, V]) {
	for i, v := range l {
		consumer(i, v)
	}
}
