package set

import "github.com/gvaligiani/algo/dict"

func NoneOf[T comparable](items map[T]struct{}, predicate func(T) bool) bool {
	return dict.NoneOf[T,struct{}](items,func (item T, _ struct{}) bool { return predicate(item) })
}