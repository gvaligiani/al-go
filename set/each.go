package set

import "github.com/gvaligiani/al.go/dict"

func Each[T comparable, S ~map[T]struct{}](items S, consumer Consumer[T]) {
	dict.Each(items, func(item T, _ struct{}) { consumer(item) })
}
