package set

import "github.com/gvaligiani/algo/dict"

func FindIfNot[T comparable](items map[T]struct{}, predicate func(T) bool) (T, bool) {
	item, _, found := dict.FindIfNot[T,struct{}](items,func (item T, _ struct{}) bool { return predicate(item) })
	return item, found
}