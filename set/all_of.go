package set

import "github.com/gvaligiani/algo/dict"

func AllOf[T comparable](items map[T]struct{}, predicate func(T) bool) bool {
	return dict.AllOf[T,struct{}](items,func (item T, _ struct{}) bool { return predicate(item) })
}