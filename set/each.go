package set

import "github.com/gvaligiani/algo/dict"

func Each[T comparable](items map[T]struct{}, consumer func(T)) {
	dict.Each[T,struct{}](items,func (item T, _ struct{}) { consumer(item) })
}