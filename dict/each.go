package dict

import "github.com/gvaligiani/al.go/util"

func Each[K comparable, V any, D ~map[K]V](d D, consumer util.Consumer[V]) {
	EachKey(d, util.ConsumeOnSecondArg[K](consumer))
}

func EachKey[K comparable, V any, D ~map[K]V](d D, consumer util.BiConsumer[K, V]) {
	for k, v := range d {
		consumer(k, v)
	}
}
